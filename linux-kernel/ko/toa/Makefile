# Makefile for building a simple kernel module

# If KERNELDIR is not set, use the default path
KERNELDIR ?= /lib/modules/$(shell uname -r)/build

# The name of the module
MODULE_NAME := toa

# The source file for the module
obj-m := $(MODULE_NAME).o

# Default target
all:
	$(MAKE) -C $(KERNELDIR) M=$(PWD) modules

# Clean target
clean:
	$(MAKE) -C $(KERNELDIR) M=$(PWD) clean

# Install target
install:
	$(MAKE) -C $(KERNELDIR) M=$(PWD) modules_install

# Uninstall target
uninstall:
	@echo "Removing module $(MODULE_NAME)"
	@sudo rmmod $(MODULE_NAME)
	@echo "Deleting module file"
	@sudo rm -f /lib/modules/$(shell uname -r)/extra/$(MODULE_NAME).ko
	@sudo depmod -a
