/*
 *Filename: octal2str.v
 *
 *Description: convert the octal numbers of input to string(ascll)
 *
 *Author: rutk1t0r
 *
 *Date:2016.12.2
 *
 *GPL
 *
 *4example:a website need I answer a question:
 *./octal2str 127 150 141 164 047 163 040 141 040 160 157 154 171 147 157 156 040 167 151 164 150 040 146 157 165 162 040 165 156 145 161 165 141 154 040 163 151 144 145 163 040 143 141 154 154 145 144 077
 *result is : What's a polygon with four unequal sides called?
 */




#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define is_octal(x) ((x>=0x30 && x<=0x37))
int octalstr2num(const char *octal)
{
	int i = 0;
	int result = 0;
	for(; i<strlen(octal); ++i)
	{
		if(!is_octal(octal[i]))
		{
			printf("error!!!found un_octal...\n");
			exit(-1);
		}
		result = result*8 + (octal[i]-0x30);
	}
	return result;

}


int main(int argc, char *argv[])
{
	int i = 1;
	char str[1024] = {0};
	if(argc < 2)
	{
		printf("Usage:\n%s [octal_nums],4example:%s 001 002 004\n", argv[0],argv[0]);
		return -1;
	}
	for(; i<argc; i++)
	{
		sprintf(&str[i-1], "%c", octalstr2num(argv[i]));
	}
	printf("result is : %s\n", str);
	
	return 0;
}
