/*
 *Filename:base32_to_hex.c
 *
 *Description:convert the base32 code to hex
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

char *base32 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567";
char *example = "SCC2WWKVWVS7EZICVDG5KBK4R4TG2BEW";
void Base32_to_Hex(char *base32str, char *hex)
{
	int i = 0;
	char tmp[6] = {0};
	for(; i < 32; i+=4)
	{
		long b3 = strchr(base32, base32str[i+0]) - base32;
		long b2 = strchr(base32, base32str[i+1]) - base32;
		long b1 = strchr(base32, base32str[i+2]) - base32;
		long b0 = strchr(base32, base32str[i+3]) - base32;
		long b =  b3 << 15 | b2 << 10 | b1 << 5 | b0;
		sprintf(tmp,"%05X",b);
		strcat(hex, tmp);
	}
}
int main(int argc, char *argv[])
{
	char example_result[44] = {0};	
	int i = 0;
	if(argc != 2 || strlen(argv[1]) != 32)
	{
		printf("Input error,Please try again!\n");
		Base32_to_Hex(example, example_result);
		printf("the example is %s %s,the result is %s ,Good luck!!!\n",argv[0], example, example_result);
		return 1;
	}
	Base32_to_Hex(argv[1], example_result);
	puts(example_result);
	return 0;
}
