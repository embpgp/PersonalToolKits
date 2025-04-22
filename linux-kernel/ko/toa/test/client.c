// client.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>

#define PORT 8080

void get_peer_ip(int sockfd) {
    struct sockaddr_in addr;
    socklen_t addr_len = sizeof(addr);
    char ip_str[INET_ADDRSTRLEN];

    if (getpeername(sockfd, (struct sockaddr *)&addr, &addr_len) == -1) {
        perror("getpeername");
        exit(EXIT_FAILURE);
    }

    inet_ntop(AF_INET, &addr.sin_addr, ip_str, sizeof(ip_str));
    printf("Connected to server IP: %s\n", ip_str);
}

int main() {
    int sock = 0;
    struct sockaddr_in serv_addr;
    char *hello = "Hello from client";

    if ((sock = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("Socket creation error");
        exit(EXIT_FAILURE);
    }

    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(PORT);

    if (inet_pton(AF_INET, "127.0.0.1", &serv_addr.sin_addr) <= 0) {
        perror("Invalid address/ Address not supported");
        close(sock);
        exit(EXIT_FAILURE);
    }

    if (connect(sock, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) < 0) {
        perror("Connection Failed");
        close(sock);
        exit(EXIT_FAILURE);
    }

    get_peer_ip(sock);

    send(sock, hello, strlen(hello), 0);
    printf("Hello message sent\n");

    close(sock);
    return 0;
}
