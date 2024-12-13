#include <iostream>
#include <cstring>
#include <unistd.h>
#include <arpa/inet.h>
#include <chrono>
#include <thread>

#define SERVER_IP "9.9.9.9"  // 目标 IP 地址
#define SERVER_PORT 12345
#define PACKET_COUNT 1000
#define SLEEP_DURATION 10
#define MAX_PACKET_SIZE 1472 // 1500 - 20 (IP header) - 8 (UDP header)

int main() {
    int sockfd;
    struct sockaddr_in serverAddr;
    char buffers[PACKET_COUNT][MAX_PACKET_SIZE]; // 二维数组，每个数据包有独立的缓冲区

    // 创建 UDP 套接字
    if ((sockfd = socket(AF_INET, SOCK_DGRAM, 0)) < 0) {
        perror("Socket creation failed");
        return -1;
    }

    // 设置服务器地址
    memset(&serverAddr, 0, sizeof(serverAddr));
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_port = htons(SERVER_PORT);
    serverAddr.sin_addr.s_addr = inet_addr(SERVER_IP);

    //记录开始时间
    struct mmsghdr msgs[PACKET_COUNT];
    struct iovec iov[PACKET_COUNT];
    //填充数据包的时候持续累积时间，超过某个时间差或者达到数据包的数量级启动sendmmsg发送数据包
    for (int i = 0; i < PACKET_COUNT; ++i) {
        snprintf(buffers[i], sizeof(buffers[i]), "Hello, UDP! Packet %d", i);
        iov[i].iov_base = buffers[i];
        iov[i].iov_len = strlen(buffers[i]);
        msgs[i].msg_hdr.msg_iov = &iov[i];
        msgs[i].msg_hdr.msg_iovlen = 1;
        msgs[i].msg_hdr.msg_name = &serverAddr;
        msgs[i].msg_hdr.msg_namelen = sizeof(serverAddr);
        msgs[i].msg_hdr.msg_control = NULL;
        msgs[i].msg_hdr.msg_controllen = 0;
        msgs[i].msg_hdr.msg_flags = 0;
    }

    for (;;) {
        int sentPackets = sendmmsg(sockfd, msgs, PACKET_COUNT, 0);
        if (sentPackets < 0) {
            perror("Sendmmsg failed");
            close(sockfd);
            return -1;
        }
        std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_DURATION));
    }

    // 关闭套接字
    close(sockfd);
    
    std::cout << "Finished sending packets." << std::endl;
    return 0;
}
