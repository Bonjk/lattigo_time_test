all:test

build:
	go build -C bfv_add
	go build -C bfv_mul
	go build -C ckks_add
	go build -C ckks_mul

clear:
	rm -f bfv_add/bfv_add bfv_mul/bfv_mul ckks_add/ckks_add ckks_mul/ckks_mul

test:build
	@echo "BFV_ADD"
	@time bfv_add/bfv_add
	@echo "BFV_MUL"
	@time bfv_mul/bfv_mul
	@echo "CKKS_ADD"
	@time ckks_add/ckks_add
	@echo "CKKS_MUL"
	@ckks_mul/ckks_mul