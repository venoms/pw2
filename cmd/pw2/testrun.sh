rm -rf git gpg; go build -v . ; ./pw2 -webinterface -logtostderr -v 9 ; rm pw2

(cd git; git log --stat ;)
