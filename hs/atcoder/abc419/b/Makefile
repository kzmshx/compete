fmt:
	ormolu --mode inplace $(shell find . -name '*.hs')

run:
	stack Main.hs

test:
	oj test --command "stack Main.hs" --directory tests --tle 2 --jobs 10
