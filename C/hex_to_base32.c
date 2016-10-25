/*
 *Filename:hex_to_base32.c
 *
 *Description:convert the hex code to base32 code
 *
 *Author:rutk1t0r
 *
 *Date:2016.8.4
 *
 *GPL
 *
 */

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <math.h>   //为了使用pow函数 linux下链接需要-lm选项将数学库引进来

void StrToHexValue(char *str, long *hex);
void Usage(char *argv[]);
int Hex2Base32(char *hex, char *str);


char example[] = "9085AB5955B565F26502A8CDD5055C8F266D0496";
int main(int argc, char *argv[])
{
	/* 在这里直接借用argv[1]传递hex编码的参数进来，假设不出问题啊
	 * hex编码为40字节,但是显示出来的时候已经成了字符串了
	 * 不能直接进行操作，需先进行转,我这里自己简单地实现了一个转化
	 * 而base32编码仅为32字节
	 */
	char str[36] = {0};
	if(argc != 2)
	{
		Usage(argv);
		return 1;
	}
	if(!Hex2Base32(argv[1], str))
	{
		Usage(argv);
		return 1;
	}
	printf("%s\n", str);
	return 0;
}

int Hex2Base32(char *hex, char *str)
{
	char *base32 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567";
	char param_table[8][6] = {0}; //这里都多留出了一个字符存储\0
	char buf[36] = {0};
	int i = 0;
	int j = 0;
	if(strlen(hex) != 40)
	{
		printf("You input is error,please try again!\n");
		return 0;
	}

	//思路是先分割，利用strncpy函数
	for(; i < 40; i+=5,++j)
	{
		//原型char *strncpy(char *dest, char *src, size_t n)
		//每次拷贝5字节的字符串即对应的十六进制值
		strncpy(param_table[j], &hex[i], 5);	
		//printf("%s\n",param_table[j]);
	}

	for(i = 0,j = 0; i<8; ++i,j+=4)
	{
		long hex = 0;   //直接传地址进去返回hex的数字值
	    StrToHexValue(param_table[i], &hex);
		//逆向操作
		long b3 = (hex >> 15) & 0x001f; 
		long b2 = (hex >> 10) & 0x001f;
		long b1 = (hex >> 5)  & 0x001f;
		long b0 = (hex >> 0)  & 0x001f;
		//直接在base32表中索引即可
		buf[j + 0] = base32[b3];
		buf[j + 1] = base32[b2];
		buf[j + 2] = base32[b1];
		buf[j + 3] = base32[b0];
	}
	strcpy(str, buf);
	return !0;
}

void StrToHexValue(char *str, long *hex)
{
	int i = 0;
	*hex = 0;
	int j = strlen(str) - 1;
	while(str[i])
	{
		int tmp = str[i] - '0';  //可以先得到差量
		if(tmp >=10 && tmp <= 16)
		{
			//标点符号
			return ;
		}
		//根据ascll码关系
		if(tmp >= 17 && tmp <= 22)
		{
			tmp -= 7;
		}
		//上述的-7是为了在这里统一运算
		if(tmp >= 0 && tmp <= 15)
		{
			*hex += tmp * pow(0x10,j);  //简单的转换，累加
		}
		++i;
		--j;
	}
}

void Usage(char *argv[])
{
	char example_result[36] = {0};
	puts(" ============================================================================================================== ");
	puts("        =     =  = = = =  =        =   =   =        = = =      = = = =   =  =       = = = =   = =       =  =    ");	
	puts("       =     =  =          =     =   =      =      =     =    =        =      =    =        =    =    =      =  ");
	puts("      =     =  =            =  =          =       = = = =    =           =        =            =           =    ");
	puts("     = = = =  = = = =        =          =        = = =      = = = =        =     = = = =     =            =     ");
	puts("    =     =  =             =  =       =         =     =    =                 =  =              =       =        ");
	puts("   =     =  =            =     =     =         =       =  =                  = =            =    =   =          ");
	puts("  =     =  = = = = =   =        =    = = = =  = = = = =  = = = = =    =  =  = = = = = =       = =     = = = =   ");
	puts(" ============================================================================================================== ");
	puts("Description:  convert the hex string to the base32 string");
	puts("");
	puts("Author:	rutk1t0r");
	puts("");
	puts("Date:	2016.08.04");
	puts("");
	puts("Copyright:	GPL");
	printf("Usage: %s HEX_STR\n", argv[0]);
	Hex2Base32(example, example_result);
	printf("example: %s %s,the result is %s\n", argv[0], example, example_result);
}
