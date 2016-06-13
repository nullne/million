package main
import (
	"regexp"
	"fmt"
)

const (
    PS1PATTERN     = `<\|(\d+)\|>`
    CTRLSEQPATTERN = `
                   \x1b[ #%()*+\-.\/]. |
                   \r |
                   (?:\x1b\[|\x9b) [ -?]* [@-~] |
                   (?:\x1b\]|\x9d) .*? (?:\x1b\\|[\a\x9c]) |
                   (?:\x1b[P^_]|[\x90\x9e\x9f]) .*? (?:\x1b\\|\x9c) |
                   \x1b.|[\x80-\x9f]
                   `
    NOBACKSPACEPATTERN = `[^\x08][\x08]`
)

var(
	reCtrl, reBackspace *regexp.Regexp
)

func filter(ori string) string {
	return reCtrl.ReplaceAllLiteralString(reBackspace.ReplaceAllLiteralString(ori, ""), "")
}

func main() {
	str := filter(`
	Script started on Thu 12 Nov 2015 06:44:43 PM CST
	^[]0;root@localhost:~^G^[[?1034hbitch > ls^M
	anaconda-ks.cfg  ^[[0m^[[38;5;34mfoo.perl^[[0m  install.log  install.log.syslog  ^[[38;5;27mmq^[[0m  nice  ^[[38;5;9mrabbitmq-server-3.5.4-1.noarch.rpm^[[0m  ^[[38;5;27mrpmbuild^[[0m  ^[[38;5;27mtmp^[[0m  typescript  ^[[38;5;9mubuntu_vim.tar^[[0m  ^[[38;5;9mwget-1.16.tar.xz^[[0m^M
	^[[m^[]0;root@localhost:~^Gbitch > pwd^M
	/root^M
	^[]0;root@localhost:~^Gbitch > fuck^M
	bash: fuck: command not found^M
	^[]0;root@localhost:~^Gbitch > exit^M
	exit^M
	`)
	fmt.Printf("%q", str)
	//expected output
	/*
	Script started on Thu 12 Nov 2015 06:44:43 PM CST
	bitch > ls
	anaconda-ks.cfg  foo.perl  install.log  install.log.syslog  mq  nice  rabbitmq-server-3.5.4-1.noarch.rpm  rpmbuild  tmp  typescript  ubuntu_vim.tar  wget-1.16.tar.xz
	bitch > pwd
	/root
	bitch > fuck
	bash: fuck: command not found
	bitch > exit
	exit
	*/
}

func init() {
	reCtrl = regexp.MustCompile(CTRLSEQPATTERN)
	reBackspace = regexp.MustCompile(NOBACKSPACEPATTERN)
}
