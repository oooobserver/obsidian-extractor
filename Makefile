.PHONY: test_file build test_dir test_dir1

test_file:
	go build -o ob && ./ob ../../../Obsidian/Database/Mysql/Basic.md

build:
	go build -o ob

test_dir: 
	go build -o ob && ./ob ../../../Obsidian/Computer\ Network


test_dir1: 
	go build -o ob && ./ob ../../../Obsidian/Database





