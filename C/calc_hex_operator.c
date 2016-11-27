/*
 *Filename:calc_hex_operator.c
 *
 *Description:simple hex operator
 *
 *Author:rutk1tor
 *
 *Date:2016.11.27
 *
 *GPL
 */

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define NUM_OPERATOR 8
char *my_operator[NUM_OPERATOR] = {
	"-add",
	"-sub",
	"-mul",
	"-div",
	"-and",
	"-or",
	"-neg",
	"-xor"
};

void Usage(char *argv[])
{
    int i = 0;
    printf("Usage:\n");
    printf("%s: [",argv[0]);
    for(; i < NUM_OPERATOR; ++i)
    {
        printf(my_operator[i]);
        printf(" | ");
    }
    printf("]\n");

}

/*将大写字母转换成小写字母*/ 
int tolower(int c) 
{ 
    if (c >= 'A' && c <= 'Z') 
    { 
        return c + 'a' - 'A'; 
    } 
    else 
    { 
        return c; 
    } 
} 

//将十六进制的字符串转换成整数 
int htoi(char s[]) 
{ 
    int i; 
    int n = 0; 
    if (s[0] == '0' && (s[1]=='x' || s[1]=='X')) //判断是否有前导0x或者0X
    { 
        i = 2; 
    } 
    else 
    { 
        i = 0; 
    } 
    for (; (s[i] >= '0' && s[i] <= '9') 
        || (s[i] >= 'a' && s[i] <= 'z') || (s[i] >='A' && s[i] <= 'Z');++i) 
    {         
        if (tolower(s[i]) > '9') 
        { 
            n = 16 * n + (10 + tolower(s[i]) - 'a'); 
        } 
        else 
        { 
            n = 16 * n + (tolower(s[i]) - '0'); 
        } 
    } 
    return n; 
} 


int calc_operator(int argc, char *argv[], int flag)
{
	int result = htoi(argv[2]);
	int i;
	if(flag == 1)
	{
		for(i = 3; i < argc; ++i)
		{
			result += htoi(argv[i]);
		}
		return result;
	}
	else if(flag == 2)
    {
        for(i = 3; i < argc; ++i)
        {
            result -= htoi(argv[i]);
        }
        return result;
    }
	else if(flag == 3)
    {
        for(i = 3; i < argc; ++i)
        {
            result *= htoi(argv[i]);
        }
        return result;
    }
	else if(flag == 4)
    {
        for(i = 3; i < argc; ++i)
        {
            result -= htoi(argv[i]);
        }
        return result;
    }
	else if(flag == 5)
    {
		
        for(i = 3; i < argc; ++i)
        {
            result &= htoi(argv[i]);
        }
        return result;
    }
	else if(flag == 6)
    {
        for(i = 3; i < argc; ++i)
        {
            result |= htoi(argv[i]);
        }
        return result;
    }
	else if(flag == 7)
    {
        return ~result;
    }
	else if(flag == 8)
    {
        for(i = 3; i < argc; ++i)
        {
            result ^= htoi(argv[i]);
        }
        return result;
    }
}

int main(int argc, char *argv[])
{
	int result = 0;
	int i;
	if(argc <= 2)
	{
		Usage(argv);
		return -1;
	}
	for(i = 0; i<NUM_OPERATOR; ++i)
	{
		if(!strcmp(argv[1], my_operator[i]))
		{
			result = calc_operator(argc, argv, i+1);
			printf("result is 0x%x\n", result);
			return 0;
		}
	}
	printf("No operator\n");
	return 1;
}



