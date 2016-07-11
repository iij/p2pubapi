#! /usr/bin/python

import os
import sys
import re
import json
import lxml.html
import urllib.request

def gettbl(root, title):
  paramtitle=root.xpath('//div[@class="section"]/h2[text()="%s"]' %(title))
  if len(paramtitle)==0: return None,None
  params=paramtitle[0].getparent()
  heads=dict([(f.get("id"), f.text_content().strip()) for f in params.xpath('.//th')])
  names=[f.text_content().strip() for f in params.xpath('.//th')]
  opt=[]
  pent=[]
  for p in params.xpath('.//td'):
    hdrid=p.get("headers").strip()
    if hdrid in heads:
      pent.append(" ".join([f.strip() for f in p.text_content().split("\n")]).strip())
      if names[-1]==heads[hdrid]:
        opt.append(pent)
        pent=[]
  return names, opt

def typefix(s):
  tbl={"String":"string", "Stinrg":"string", "Array":"[]string"}
  if s in tbl:
    return tbl[s]
  return s

def snake2camel(s):
  return "".join([x.capitalize() for x in s.split("-")])

def apiname2(d):
  namemap={
    "仮想サーバ": "VM",
    "追加申込": "Add",
    "一覧": "List",
    "情報": "",
    "契約状態": "Contract",
    "取得": "Get",
    "リソース状態": "Resource",
    "ラベル": "Label",
    "設定": "Set",
    "起動": "Power",
    "停止": "Power",
    "再起動": "Power",
    "ブートデバイス": "BootDevice",
    "データデバイス": "DataDevice",
    "ストレージ": "Storage",
    "切断": "Disconnect",
    "接続": "Connect",
    "グローバルIPアドレス": "GlobalAddress",
    "グローバルIPv6": "GlobalIPv6",
    "グローバルIPアドレス/V": "GlobalAddressV",
    "割り当て": "Allocate",
    "解放": "Release",
    "FW+LB": "FwLb",
    "追加ストレージ": "Storage",
    "解約申込": "Cancel",
    "プライベート": "Private",
    "ネットワーク": "Network",
    "ネットワーク/V": "NetworkV",
    "カスタム": "Custom",
    "イメージ": "Image",
    "P2（パブリックリソース）": "P2PUB",
    "契約": "Contract",
    "（SGSA向け）": "ForSGSA",
    "（SA向け）": "ForSA",
    "公開鍵": "PublicKey",
    "追加": "Add",
    "逆引き": "Reverse",
    "ドメイン": "Domain",
    "パスワード": "Password",
    "冗長構成": "Redundancy",
    "変更申込": "Change",
    "システム": "System",
    "アーカイブ": "Archive",
    "初期化": "Initialize",
    "品目": "Item",
    "管理画面": "ControlPanel",
    "アクセス制限": "ACL",
    "削除": "Delete",
    "アドレス数": "AddressNum",
    "リストア": "Restore",
    "サービス": "Service",
    "コード": "Code",
    "アカウント": "Account",
    "無効": "Disable",
    "有効": "Enable",
    "（グローバルIPアドレス）": "ForGlobal",
    "（仮想サーバ）": "ForVM",
    "割当": "Allocate",
    "名称": "Name",
    "変更": "Change",
    "セットアップ": "Setup",
    "作成": "Create",
    "名": "Name",
    "サイズ": "Size",
  }
  for k in sorted(namemap, key=lambda k: len(k), reverse=True):
    d=d.replace(k,namemap[k])
  return d

def apiname(d):
  mtd=d["メソッド"]
  uris=d["URI"].split(" ")
  name=""
  for m in uris[0].split("/"):
    if m.startswith(":"): continue
    name+=snake2camel(os.path.splitext(m)[0])
  return name+mtd.capitalize()

def async(d):
  return d["実行"]

def template(d):
  uris=d["URI"].split(" ")
  res=[]
  keys=[]
  uri=uris[0].split("?")[0]
  nuri=uri
  for m in re.findall(":[a-zA-Z]+", uri):
    nuri=nuri.replace(m, "{{."+m.replace(":", "")+"}}")
    keys.append(m.strip(":"))
  return nuri, keys

def method(d):
  return d["メソッド"]

def open_or_download(fname, baseurl):
  if not os.path.isdir("cache"):
    os.mkdir("cache")
  if not os.path.exists(os.path.join("cache", fname)):
    urllib.request.urlretrieve(os.path.join(baseurl, fname), os.path.join("cache", fname))
  return open(os.path.join("cache", fname))

def do1(fname, baseurl):
  fp=open_or_download(fname, baseurl)
  res={}
  root=lxml.html.parse(fp)

  res["title"]=root.xpath("//title")[0].text
  head, body = gettbl(root, "API情報")
  if head==None: return None
  key=dict(zip(head, body[0]))
  #res["name"]=apiname(key)
  res["name"]=apiname2(res["title"])
  assert(re.match("^[a-zA-Z0-9]*$", res["name"]))
  res["template"], tmpl_args=template(key)
  res["method"]=method(key)
  res["async"]=async(key)
  # print("api-info", res)

  res["template_arg"]=[]
  res["query_arg"]=[]
  res["json_arg"]=[]
  res["json_resp"]=[]

  head, body = gettbl(root, "リクエストパラメータ")
  if head[0]=="":
    head[0]="type"
  for b in body:
    if b[0]=="":
      b[0]=preb
    preb=b[0]
    key=dict(zip(head, b))
    optional = key["必須"]!="○"
    val=[key["パラメータ"].split(" ")[0].split("または")[0], optional, key["意味"], key["値"]]
    if key["type"]=="クエリストリング" or key["パラメータ"].endswith("Index") or key["パラメータ"].endswith("Count"):
       res["query_arg"].append(val)
    elif key["type"]=="URL":
       res["template_arg"].append(val)
    elif key["type"]=="ボディ":
       res["json_arg"].append(val)
  for t in tmpl_args:
    if t not in [x[0] for x in res["template_arg"]]:
      res["template_arg"].append([t, True, t, ""])
  # print("req-info", res)

  head, body = gettbl(root, "レスポンス")
  for b in body:
    key=dict(zip(head, b))
    val=[key["フィールド"], key["タイプ"], key["意味"], key["値"]]
    res["json_resp"].append(val)
  # print("resp-info", res)

  return res

def getindex(fp):
  res={}
  root=lxml.html.parse(fp)
  apilist_links=set(filter(lambda f: f.startswith("5"), [x.attrib['href'] for x in root.xpath('//a')]))
  return apilist_links

def merge(a, b):
  res=a[:]
  for i in b:
    if i[0] not in map(lambda f: f[0], res):
      res.append(i)
    elif i[3] and i[3]!="":
      sample_add=i[3]
      for x in res:
        if x[0]==i[0]:
          if x[3] and x[3]!="":
            x[3]+=","+i[3]
          else:
            x[3]=i[3]
  return res

final={}

base="http://manual.iij.jp/p2/pubapi/"

if len(sys.argv)==1:
  start="58436087.html"
  names=getindex(open_or_download(start, base))
else:
  names=sys.argv[1:]

fnmap={}

for i in names:
  x=do1(i, base)
  if x!=None:
    name=x["name"]
    if name in final:
      print("duplicate entry:", i, name)
      # merge
      oldval=final[name]
      x["template_arg"]=merge(x["template_arg"], oldval["template_arg"])
      x["query_arg"]=merge(x["query_arg"], oldval["query_arg"])
      x["json_arg"]=merge(x["json_arg"], oldval["json_arg"])
      x["json_resp"]=merge(x["json_resp"], oldval["json_resp"])
    final[name]=x
    if name not in fnmap:
      fnmap[name]=[i]
    else:
      fnmap[name].append(i)

def show_jsonmap(prefix, args, idx, tidx, tblmap, fp):
  for k,v in tblmap.items():
    if prefix:
      fullname=prefix+"."+k
    else:
      fullname=k
    entlist=list(filter(lambda f: f[0]==fullname, args))
    if len(entlist)==0:
      entry=["",False,"",""]
    else:
      entry=entlist[0]
    print("name", fullname, "entry", entry)
    descr=entry[2]
    sample=""
    if entry[3]!="":
      sample="("+"/".join(set([x.strip().strip('"') for x in entry[3].split(",")]))+")"
    if entry[1]:
      req="`json:\",omitempty\"`"
    else:
      req=""
    if v==True:  # tail
      print("  {k} string {req} // {descr}{sample}".format(*entry, k=k, req=req, descr=descr, sample=sample), file=fp)
    else:
      if k.endswith("List"):
        ary="[]"
      else:
        ary=""
      print(" {k} {ary}struct{{".format(k=k, ary=ary), file=fp)
      show_jsonmap(fullname, args, idx, tidx, v, fp)
      print("} `json:\",omitempty\"`", file=fp)

def json_arg(args, idx=None, tidx=None, fp=sys.stdout):
  print("json-arg:", args, idx, tidx)
  if len(args)==0: return
  # make struct
  tblmap={}
  for i,t in [(a[0],a[1]) for a in args]:
    #print("i", i, t)
    if t in ("Array", "Object") or i.endswith("List"): continue
    cur=tblmap
    for k in i.split("."):
      tailp=i.endswith(k)
      #print("k", k, tailp)
      if k not in cur:
        if tailp:
          cur[k]=True
        else:
          cur[k]={}
      cur=cur[k]
  print("tblmap", tblmap)
  flg = (args[0][0].endswith("List") and args[0][0].find(".")==-1)
  print("first", args[0][0])
  if flg:
      print("{0} []struct{{".format(args[0][0]), file=fp)
  show_jsonmap(None, args, idx, tidx, tblmap, fp)
  if flg:
      print("}", file=fp)

#print("final", final)

# output
for k,v in final.items():
  print("output:", k)
  with open("%s.go" %(k), "w") as fp:
  #with sys.stdout as fp:
    print("""package protocol
import (
"reflect"
)
// {name} {title} ({async})""".format(**v), file=fp)
    for fn in fnmap[v["name"]]:
      print("//  {baseurl}{item}".format(baseurl=base, item=fn), file=fp)
    print("type {name} struct {{".format(**v), file=fp)
    for i in v["template_arg"]:
      sample=""
      if i[3] and i[3]!="":
        sample="("+i[3]+")"
      print("  {0} string `json:\"-\"` // {2}{sample}".format(*i, sample=sample), file=fp)
    for i in v["query_arg"]:
      sample=""
      if i[3] and i[3]!="":
        sample="("+i[3]+")"
      print("  {0} string `json:\"-\" p2pub:\",query\"`  // {2}{sample}".format(*i, sample=sample), file=fp)
    json_arg(v["json_arg"], idx=1, fp=fp)
    print("}", file=fp)
    print("""// URI {template}
func (t {name}) URI() string {{
  return \"{template}\"
}}
// APIName {name}
func (t {name}) APIName() string {{
  return \"{name}\"
}}
// Method {method}
func (t {name}) Method() string {{
  return "{method}"
}}
// Document {baseurl}{fname}
func (t {name}) Document() string {{
  return "{baseurl}{fname}"
}}
// JPName {title}
func (t {name}) JPName() string {{
  return "{title}"
}}
func init(){{
  APIlist=append(APIlist, {name}{{}})
  TypeMap["{name}"]=reflect.TypeOf({name}{{}})
}}""".format(**v, baseurl=base, fname=fnmap[v["name"]][0]), file=fp)
    print("// {name}Response {title}のレスポンス".format(**v), file=fp)
    print("type {name}Response struct {{".format(**v), file=fp)
    print("*CommonResponse", file=fp)
    json_arg(v["json_resp"], tidx=1, fp=fp)
    print("}", file=fp)

os.system("go fmt .")
