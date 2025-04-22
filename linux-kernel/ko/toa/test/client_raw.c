// client_raw.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <netinet/tcp.h>
#include <netinet/ip.h>

#define PORT 8080

// Pseudo header needed for TCP checksum calculation
struct pseudo_header {
    u_int32_t source_address;
    u_int32_t dest_address;
    u_int8_t placeholder;
    u_int8_t protocol;
    u_int16_t tcp_length;
};

// TCP checksum calculation function
unsigned short checksum(void *b, int len) {
    unsigned short *buf = b;
    unsigned int sum = 0;
    unsigned short result;

    for (sum = 0; len > 1; len -= 2)
        sum += *buf++;
    if (len == 1)
        sum += *(unsigned char *)buf;
    sum = (sum >> 16) + (sum & 0xFFFF);
    sum += (sum >> 16);
    result = ~sum;
    return result;
}

void create_ip_header(struct iphdr *iph, struct sockaddr_in *serv_addr) {
    iph->ihl = 5;
    iph->version = 4;
    iph->tos = 0;
    iph->tot_len = sizeof(struct iphdr) + sizeof(struct tcphdr);
    iph->id = htonl(54321);
    iph->frag_off = 0;
    iph->ttl = 255;
    iph->protocol = IPPROTO_TCP;
    iph->check = 0;
    iph->saddr = inet_addr("127.0.0.1");
    iph->daddr = serv_addr->sin_addr.s_addr;

    // IP checksum
    iph->check = checksum((unsigned short *)iph, iph->tot_len);
}

void create_tcp_header(struct tcphdr *tcph, int src_port, int dest_port, int seq, int ack_seq, int syn, int ack) {
    tcph->source = htons(src_port);
    tcph->dest = htons(dest_port);
    tcph->seq = htonl(seq);
    tcph->ack_seq = htonl(ack_seq);
    tcph->doff = 5; // TCP header size
    tcph->fin = 0;
    tcph->syn = syn;
    tcph->rst = 0;
    tcph->psh = 0;
    tcph->ack = ack;
    tcph->urg = 0;
    tcph->window = htons(5840); /* maximum allowed window size */
    tcph->check = 0; // Leave checksum 0 now, filled later by pseudo header
    tcph->urg_ptr = 0;
}

void calculate_tcp_checksum(struct tcphdr *tcph, struct pseudo_header *psh) {
    int psize = sizeof(struct pseudo_header) + sizeof(struct tcphdr);
    char *pseudogram = malloc(psize);

    memcpy(pseudogram, (char *)psh, sizeof(struct pseudo_header));
    memcpy(pseudogram + sizeof(struct pseudo_header), tcph, sizeof(struct tcphdr));

    tcph->check = checksum((unsigned short *)pseudogram, psize);

    free(pseudogram);
}

int main() {
    int sock;
    struct sockaddr_in serv_addr;
    char packet[4096];
    char recv[4096];
    struct iphdr *iph = (struct iphdr *)packet;
    //struct tcphdr *tcph = (struct tcphdr *)(packet)
    struct tcphdr *tcph = (struct tcphdr *)(packet+sizeof(struct iphdr));
    struct pseudo_header psh;

    // Create raw socket
    if ((sock = socket(AF_INET, SOCK_RAW, IPPROTO_TCP)) < 0) {
        perror("Socket creation error");
        exit(EXIT_FAILURE);
    }

    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(PORT);
    serv_addr.sin_addr.s_addr = inet_addr("127.0.0.1");

    memset(packet, 0, 4096);
    int one = 1;
    if (setsockopt(sock, IPPROTO_IP, IP_HDRINCL, &one, sizeof(one)) < 0) {
        perror("setsockopt");
        exit(EXIT_FAILURE);
    }
    // Fill in the IP Header
    create_ip_header(iph, &serv_addr);

    // Fill in the TCP Header
    create_tcp_header(tcph, 12345, PORT, 0, 0, 1, 0);
        // Now the TCP checksum
    psh.source_address = inet_addr("127.0.0.1");
    psh.dest_address = serv_addr.sin_addr.s_addr;
    psh.placeholder = 0;
    psh.protocol = IPPROTO_TCP;
    psh.tcp_length = htons(sizeof(struct tcphdr));

    calculate_tcp_checksum(tcph, &psh);
    //int send_len = sizeof(struct tcphdr);
    int send_len = iph->tot_len;
    // Send SYN packet
    if (sendto(sock, packet, send_len, 0, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) < 0) {
        perror("sendto failed");
        exit(EXIT_FAILURE);
    } else {
        printf("SYN Packet Sent\n");
    }

    // Receive SYN-ACK packet
    socklen_t serv_addr_len = sizeof(serv_addr);
    if (recvfrom(sock, recv, 4096, 0, (struct sockaddr *)&serv_addr, &serv_addr_len) < 0) {
        perror("recvfrom failed");
        exit(EXIT_FAILURE);
    } else {
        printf("SYN-ACK Packet Received\n");
    }

    // Extract TCP header from received packet
    struct tcphdr *recv_tcph = (struct tcphdr *)(recv + sizeof(struct iphdr));

    // Check if the received packet is a SYN-ACK packet
    if (recv_tcph->syn == 1 && recv_tcph->ack == 1) {
        // Create ACK packet
        memset(packet, 0, sizeof(packet));
        create_tcp_header(tcph, 12345, PORT, ntohl(recv_tcph->ack_seq), ntohl(recv_tcph->seq) + 1, 0, 1);

        // Calculate TCP checksum
        calculate_tcp_checksum(tcph, &psh);

        // Send the ACK packet
        if (sendto(sock, packet, send_len, 0, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) < 0) {
            perror("Sendto failed");
            exit(EXIT_FAILURE);
        }

        printf("ACK Packet Sent......\n");
    } else {
        printf("Received packet is not a SYN-ACK packet\n");
    }

    // Add data to the packet
    char *data = packet + sizeof(struct tcphdr);
    const char *msg = "Hello, Server!";
    strcpy(data, msg);
    send_len = sizeof(struct iphdr) + sizeof(struct tcphdr) + strlen(msg);

    // Recalculate TCP checksum
    psh.tcp_length = htons(sizeof(struct tcphdr) + strlen(msg));
    calculate_tcp_checksum(tcph, &psh);

    // Send data packet
    if (sendto(sock, packet, send_len, 0, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) < 0) {
        perror("sendto failed");
        exit(EXIT_FAILURE);
    } else {
        printf("Data Packet Sent\n");
    }

    close(sock);
    return 0;
}

