package main

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"os"
)

func compileLLVM(functionList []*Function) error {
	//testLLVM1()
	err := testLLVM2(functionList)
	return err
}

func testLLVM2(listFunction []*Function) error {
	m := ir.NewModule()

	i32 := types.I32
	zero := constant.NewInt(i32, 0)

	main := m.NewFunc("main2", i32)

	for _, function := range listFunction {
		var funct *ir.Func
		if function.Name == "main" {
			funct = main
		} else {
			funct = m.NewFunc(function.Name, i32)
		}
		entry := funct.NewBlock("")
		for _, instruction := range function.Instruction {
			if instruction.Code == INSTRUCTION_AFFECTATION {
				val := instruction.Valeur
				var c *constant.Int
				if val.code == EXPR_CODE_INT {
					c = constant.NewInt(i32, int64(val.valeurInt))
				} else {
					c = constant.NewInt(i32, 0)
				}
				tmp := entry.NewAlloca(i32)
				//tmp := entry.NewAdd(tmp2, c)
				entry.NewStore(c, tmp)
			}
		}
		entry.NewRet(zero)
	}
	//pretty.Println(m)
	f, err := os.Create("./tmp/test1.ll")
	if err != nil {
		return err
	}
	defer func(f *os.File) error {
		err := f.Close()
		if err != nil {
			return err
		}
		return nil
	}(f)
	_, err = fmt.Fprintln(f, m)
	if err != nil {
		return err
	}

	err = writeMain()
	if err != nil {
		return err
	}

	return nil
}

func writeMain() error {
	f, err := os.Create("./tmp/main.c")
	if err != nil {
		return err
	}
	defer func(f *os.File) error {
		err := f.Close()
		if err != nil {
			return err
		}
		return nil
	}(f)
	_, err = fmt.Fprintf(f, "#include<stdio.h>\n\nint main2();\n\nint main()\n{\n\tprintf(\"coucou\\n\");\n\tprintf(\"n=%%d\\n\",main2());\n\treturn 0;\n}\n")
	if err != nil {
		return err
	}
	return nil
}

func testLLVM1() {
	i32 := types.I32
	zero := constant.NewInt(i32, 0)
	a := constant.NewInt(i32, 0x15A4E35) // multiplier of the PRNG.
	c := constant.NewInt(i32, 1)         // increment of the PRNG.

	// Create a new LLVM IR module.
	m := ir.NewModule()

	// Create an external function declaration and append it to the module.
	//
	//    int abs(int x);
	abs := m.NewFunc("abs", i32, ir.NewParam("x", i32))

	// Create a global variable definition and append it to the module.
	//
	//    int seed = 0;
	seed := m.NewGlobalDef("seed", zero)

	// Create a function definition and append it to the module.
	//
	//    int rand(void) { ... }
	rand := m.NewFunc("rand", i32)

	// Create an unnamed entry basic block and append it to the `rand` function.
	entry := rand.NewBlock("")

	// Create instructions and append them to the entry basic block.
	tmp1 := entry.NewLoad(i32, seed)
	tmp2 := entry.NewMul(tmp1, a)
	tmp3 := entry.NewAdd(tmp2, c)
	entry.NewStore(tmp3, seed)
	tmp4 := entry.NewCall(abs, tmp3)
	entry.NewRet(tmp4)

	// Print the LLVM IR assembly of the module.
	fmt.Println(m)
}
