run:
	./run.sh

cgen:
	c-for-go --ccincl --ccdefs --nostamp gogc.yml

clean:
	rm -rf ./libgogc.a
