#include <stdio.h>
#include <math.h>

// 计算三角形的周长
double circum(double a, double b, double c)
{
  return a + b + c;
}

// 计算三角形的面积（使用海伦公式）
double area(double a, double b, double c)
{
  double s = (a + b + c) / 2.0;
  return sqrt(s * (s - a) * (s - b) * (s - c));
}

int main()
{
  double a, b, c;

  // 获取用户输入
  printf("a,b,c:\n");
  printf("a = ");
  scanf("%lf", &a);
  printf("b = ");
  scanf("%lf", &b);
  printf("c = ");
  scanf("%lf", &c);
  if (a + b > c && a + c > b && b + c > a)
  {
    // 计算周长和面积
    double perimeter = circum(a, b, c);
    double areaNew = area(a, b, c);

    // 输出结果
    printf("circum: %.2f\n", perimeter);
    printf("area: %.2f\n", areaNew);
  }
  else
  {
    printf("error\n");
  }

  return 0;
}