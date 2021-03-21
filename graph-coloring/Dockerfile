FROM ubuntu:rolling

RUN apt-get update && apt-get install -y git python3 python3-pip wget && pip3 install ansible
RUN git clone https://github.com/frankdressel/blog.git /tmp/bloginstall && cd /tmp/bloginstall/ansible && ansible-playbook -i "localhost," -c local vim.yml
RUN sed -i -e "s/call vundle#end()/Plugin 'fatih\/vim-go'\\ncall vundle#end()/g" /root/.vimrc && sed -i '1s/^/set encoding=utf-8\n/' /root/.vimrc && vim +PluginInstall +qall

RUN wget https://golang.org/dl/go1.16.2.linux-amd64.tar.gz -O /usr/local/go1.16.2.linux-amd64.tar.gz && cd /usr/local && tar -xzf go1.16.2.linux-amd64.tar.gz
ENV PATH "$PATH:/usr/local/go/bin"
