#include <stdio.h>


int gcd(int p,int q){
	if(q == 0) return p;
	int r = p % q;
	return gcd(q,r);
}

int main(){

	int a = 0,b = 0;
	int result = 0;

	while (scanf("%d %d",&a,&b) != EOF){

		result = gcd(a,b);
		printf("%d\n",result);

	}

	return 0;
}
