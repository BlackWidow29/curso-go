package main

import(
	"fmt"
)

func main()  {
	fmt.Println("INTEIROS SEM SINAL");

	//byte também só vão até 255
	var u1 uint8 = 255;
	var u2 uint16 = 64000;
	var u3 uint32 = 4000000;
	var u4 uint64 = 900000000000;
	fmt.Println(u1);
	fmt.Println(u2);
	fmt.Println(u3);
	fmt.Println(u4);

	fmt.Println("INTEIROS COM SINAL");
	 var i1 int8 = -127;
	 fmt.Println(i1);
	 var i2 int16 = -10000;
	 fmt.Println(i2);
	 var i3 rune = -1000001; //int32
	 fmt.Println(i3);
	 var i4 int64 = -100000002;
	 fmt.Println(i4);

	 fmt.Println("PONTOS FLUTUANTES");
	 var f1 float32 = 123444.4;
	 fmt.Println(f1);
	 var f2 float64 = 1111.2222222;
	 fmt.Println(f2);
	 var f3 complex64 = 111022222222;
	 fmt.Println(f3);
	 var f4 complex128 = 1111111111111111111111111111111111111111111111111111.222222222222222222222222222222222;
	 fmt.Println(f4);

	 fmt.Println("STRINGS");
	 var s1 string = "TreinaWeb";
	 fmt.Println(s1);
	 var s2 string = `TREINA WEB
	 CURSOS`;
	 fmt.Println(s2);

	 fmt.Println("BOOLEANOS");
	 var bool1 bool = true;
	 fmt.Println(bool1);

	 fmt.Println("MULTIPLAS DECLARAÇÕES")
	 var nome, sobrenome string = "TREINAWEB", "CURSOS";

	 fmt.Println(nome + " " + sobrenome);

	 var(
		 sexo string = "F"
		 idade int8 = 20
	 );

	 fmt.Println(sexo);
	 fmt.Println(idade);

}