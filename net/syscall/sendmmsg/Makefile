# Makefile for compiling and running sendto

# 编译器
CXX = g++

# 编译器选项
CXXFLAGS = -std=c++11 -Wall -g

# 目标可执行文件
TARGET = sendmmsg

# 源文件
SRCS = sendmmsg.cpp

# 目标文件
OBJS = $(SRCS:.cpp=.o)

# 默认目标
all: $(TARGET)

# 链接目标文件生成可执行文件
$(TARGET): $(OBJS)
	$(CXX) $(CXXFLAGS) -o $(TARGET) $(OBJS)

# 编译源文件生成目标文件
%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -o $@ 

# 运行可执行文件
run: $(TARGET)
	./$(TARGET)

# 清理生成的文件
clean:
	rm -f $(TARGET) $(OBJS)

.PHONY: all run clean
