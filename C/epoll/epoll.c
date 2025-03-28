#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/epoll.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <fcntl.h>

#define MAX_EVENTS 10000
#define BUFFER_SIZE 4096
#define HTTP_RESPONSE "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello Epoll World!\n"

// 设置文件描述符为非阻塞模式
void set_nonblocking(int fd) {
    int flags = fcntl(fd, F_GETFL, 0);
    fcntl(fd, F_SETFL, flags | O_NONBLOCK);
}

int main() {
    int listen_fd, epoll_fd;
    struct sockaddr_in addr;
    struct epoll_event ev, events[MAX_EVENTS];
    
    // 创建监听socket
    listen_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (listen_fd == -1) {
        perror("socket");
        exit(EXIT_FAILURE);
    }

    // 绑定端口
    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = htonl(INADDR_ANY);
    addr.sin_port = htons(8080);

    if (bind(listen_fd, (struct sockaddr*)&addr, sizeof(addr)) == -1) {
        perror("bind");
        close(listen_fd);
        exit(EXIT_FAILURE);
    }

    // 设置非阻塞模式
    set_nonblocking(listen_fd);

    // 开始监听
    if (listen(listen_fd, SOMAXCONN) == -1) {
        perror("listen");
        close(listen_fd);
        exit(EXIT_FAILURE);
    }

    // 创建epoll实例
    epoll_fd = epoll_create1(0);
    if (epoll_fd == -1) {
        perror("epoll_create1");
        close(listen_fd);
        exit(EXIT_FAILURE);
    }

    // 注册监听socket到epoll
    ev.events = EPOLLIN | EPOLLET; // 边缘触发模式
    ev.data.fd = listen_fd;
    if (epoll_ctl(epoll_fd, EPOLL_CTL_ADD, listen_fd, &ev) == -1) {
        perror("epoll_ctl");
        close(listen_fd);
        close(epoll_fd);
        exit(EXIT_FAILURE);
    }

    printf("Epoll HTTP Server running on port 8080...\n");

    while (1) {
        int nfds = epoll_wait(epoll_fd, events, MAX_EVENTS, -1);
        if (nfds == -1) {
            perror("epoll_wait");
            break;
        }

        for (int i = 0; i < nfds; ++i) {
            // 处理新连接
            if (events[i].data.fd == listen_fd) {
                struct sockaddr_in client_addr;
                socklen_t client_len = sizeof(client_addr);
                int conn_fd = accept(listen_fd, 
                                   (struct sockaddr*)&client_addr,
                                   &client_len);
                if (conn_fd == -1) {
                    perror("accept");
                    continue;
                }

                set_nonblocking(conn_fd);
                ev.events = EPOLLIN | EPOLLET;
                ev.data.fd = conn_fd;
                if (epoll_ctl(epoll_fd, EPOLL_CTL_ADD, conn_fd, &ev) == -1) {
                    perror("epoll_ctl");
                    close(conn_fd);
                }
            } 
            // 处理客户端请求
            else if (events[i].events & EPOLLIN) {
                char buffer[BUFFER_SIZE];
                ssize_t bytes_read;
                int client_fd = events[i].data.fd;

                // 读取请求（简单处理）
                bytes_read = read(client_fd, buffer, BUFFER_SIZE);
                if (bytes_read > 0) {
                    // 发送HTTP响应
                    write(client_fd, HTTP_RESPONSE, strlen(HTTP_RESPONSE));
                }

                // 关闭连接（保持短连接）
                close(client_fd);
            }
        }
    }

    close(listen_fd);
    close(epoll_fd);
    return 0;
}
