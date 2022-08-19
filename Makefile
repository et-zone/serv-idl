ADDDIRGO =$(shell mkdir -p gen_go)
MAKEGO =$(shell protoc --micro_out=./gen_go --go_out=./gen_go *.proto)

ADDDIRJS =$(shell mkdir -p gen_js)
MAKEJS =$(shell protoc --js_out=./gen_js *.proto)

ADDDIRJAVA =$(shell mkdir -p gen_java)
MAKEJAVA =$(shell protoc --java_out=./gen_java *.proto)

CLESNGO =$(shell rm -rf gen_go)
CLESNJS =$(shell rm -rf gen_js)
CLESNJAVA =$(shell rm -rf gen_java)

clean:
	@echo -n $(CLESNGO)
#	@echo -n $(CLESNJS)
#	@echo -n $(CLESNJAVA)
all:
	@echo -n $(ADDDIRGO)
	@echo -n $(MAKEGO)
#	@echo -n $(ADDDIRJS)
#	@echo -n $(MAKEJS)
#	@echo -n $(ADDDIRJAVA)
#	@echo -n $(MAKEJAVA)