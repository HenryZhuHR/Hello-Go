import{_ as o,r as t,o as d,c as i,a as e,b as a,d as l,e as s}from"./app-L6S3ynmI.js";const r={},u=s(`<h1 id="hello-vuepress" tabindex="-1"><a class="header-anchor" href="#hello-vuepress" aria-hidden="true">#</a> Hello VuePress</h1><h2 id="开发环境" tabindex="-1"><a class="header-anchor" href="#开发环境" aria-hidden="true">#</a> 开发环境</h2><p>安装 Go 后，配置 GOPATH， <code>export GOPATH=&quot;$HOME/project/goproject&quot;</code> ，开发时，代码存放在 <code>$GOPATH/src</code> 下， 工程经过 <code>go build</code>、<code>go install</code> 或<code>go get</code> 等指令后，会将下载的第三方包源代码文件放在 <code>$GOPATH/src</code> 目录下， 产生的二进制可执行文件放在 <code>$GOPATH/bin</code> 目录下，生成的中间缓存文件会被保存在 <code>$GOPATH/pkg</code> 下</p><p>配置代理</p><div class="language-bash line-numbers-mode" data-ext="sh"><pre class="language-bash"><code><span class="token builtin class-name">export</span> <span class="token assign-left variable">GOPROXY</span><span class="token operator">=</span><span class="token string">&quot;https://goproxy.cn,direct&quot;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="go-modules-包管理" tabindex="-1"><a class="header-anchor" href="#go-modules-包管理" aria-hidden="true">#</a> Go Modules 包管理</h2>`,6),c={href:"https://go.dev/blog/using-go-modules",target:"_blank",rel:"noopener noreferrer"},v=s(`<p>创建 modules 的命令</p><div class="language-bash line-numbers-mode" data-ext="sh"><pre class="language-bash"><code>go mod init github.com/<span class="token operator">&lt;</span>username<span class="token operator">&gt;</span>/<span class="token operator">&lt;</span>projectname<span class="token operator">&gt;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>vscode 还有一项很强大的功能就是断点调试,结合 delve 可以很好的进行 Go 代码调试</p><div class="language-bash line-numbers-mode" data-ext="sh"><pre class="language-bash"><code>go get <span class="token parameter variable">-v</span> <span class="token parameter variable">-u</span> github.com/peterh/liner github.com/derekparker/delve/cmd/dlv

brew <span class="token function">install</span> go-delve/delve/delve <span class="token punctuation">(</span>mac可选<span class="token punctuation">)</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>如果有问题再来一遍:</p><div class="language-bash line-numbers-mode" data-ext="sh"><pre class="language-bash"><code>go get <span class="token parameter variable">-v</span> <span class="token parameter variable">-u</span> github.com/peterh/liner github.com/derekparker/delve/cmd/dlv
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>注意:修改&quot;dlv-cert&quot;证书, 选择&quot;显示简介&quot;-&gt;&quot;信任&quot;-&gt;&quot;代码签名&quot; 修改为: 始终信任</p><p>打开首选项-工作区设置,配置launch.json:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>{
    &quot;version&quot;: &quot;0.2.0&quot;,
    &quot;configurations&quot;: [
        {
            &quot;name&quot;: &quot;main.go&quot;,
            &quot;type&quot;: &quot;go&quot;,
            &quot;request&quot;: &quot;launch&quot;,
            &quot;mode&quot;: &quot;debug&quot;,
            &quot;remotePath&quot;: &quot;&quot;,
            &quot;port&quot;: 2345,
            &quot;host&quot;: &quot;127.0.0.1&quot;,
            &quot;program&quot;: &quot;\${workspaceRoot}&quot;, // 工作空间路径
            &quot;env&quot;: {},
            &quot;args&quot;: [],
            &quot;showLog&quot;: true
        }
    ]
}
\`\`\`shell</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,9);function p(m,b){const n=t("ExternalLinkIcon");return d(),i("div",null,[u,e("p",null,[e("a",c,[a("Go Modules"),l(n)]),a(" 是 Go 1.11 版本后引入的包管理工具，它可以帮助我们更好的管理第三方包，解决了以往的 GOPATH 机制下，包管理不方便的问题，同时也解决了包版本管理的问题。")]),v])}const g=o(r,[["render",p],["__file","index.html.vue"]]);export{g as default};
