// server.c
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
    printf("Connected to client IP: %s\n", ip_str);
}

int main() {
    int server_fd, new_socket;
    struct sockaddr_in address;
    int opt = 1;
    int addrlen = sizeof(address);

    if ((server_fd = socket(AF_INET, SOCK_STREAM, 0)) == 0) {
        perror("socket failed");
        exit(EXIT_FAILURE);
    }

    if (setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR | SO_REUSEPORT, &opt, sizeof(opt))) {
        perror("setsockopt");
        close(server_fd);
        exit(EXIT_FAILURE);
    }

    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY;
    address.sin_port = htons(PORT);

    if (bind(server_fd, (struct sockaddr *)&address, sizeof(address)) < 0) {
        perror("bind failed");
        close(server_fd);
        exit(EXIT_FAILURE);
    }

    if (listen(server_fd, 3) < 0) {
        perror("listen");
        close(server_fd);
        exit(EXIT_FAILURE);
    }

    printf("Server listening on port %d\n", PORT);

    if ((new_socket = accept(server_fd, (struct sockaddr *)&address, (socklen_t*)&addrlen)) < 0) {
        perror("accept");
        close(server_fd);
        exit(EXIT_FAILURE);
    }

    get_peer_ip(new_socket);

    char buffer[1024] = {0};
    read(new_socket, buffer, 1024);
    printf("Message from client: %s\n", buffer);

    close(new_socket);
    close(server_fd);
    return 0;
}
