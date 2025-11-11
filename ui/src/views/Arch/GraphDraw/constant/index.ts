export interface Node {
  type: string
  class: string
  text: string
}

export const graphNodes: Node[] = [{ type: 'archGraph', class: 'arch_graph', text: '架构图' }]

// 常用组件节点列表
export const usedNodes: Node[] = [
  { type: 'ecs', class: 'ecs', text: 'ecs' },
  { type: 'slb', class: 'slb', text: 'slb' },
  { type: 'kong', class: 'kong', text: 'kong' },
  { type: 'nginx', class: 'nginx', text: 'nginx' },
  { type: 'mq', class: 'mq', text: 'mq' },
  { type: 'oss', class: 'oss', text: 'oss' },
  { type: 'waf', class: 'waf', text: 'waf' },
  { type: 'rds', class: 'rds', text: 'rds' },
  { type: 'drds', class: 'drds', text: 'drds' },
  { type: 'redis', class: 'redis', text: 'redis' },
  { type: 'javaApp', class: 'java_app', text: 'java应用' },
  { type: 'domain', class: 'domain', text: '域名' },
  { type: 'switch', class: 'switch', text: '交换机' },
  { type: 'router', class: 'router', text: '路由器' },
  { type: 'firewall', class: 'firewall', text: '防火墙' },
  { type: 'physicalMachine', class: 'physical_machine', text: '物理机' },
  { type: 'gbase', class: 'gbase', text: 'gbase' },
  { type: 'hbase', class: 'hbase', text: 'hbase' },
  { type: 'oracle', class: 'oracle', text: 'oracle' },
  { type: 'browser', class: 'browser', text: '浏览器' },
  { type: 'es', class: 'es', text: 'elasticsearch' },
  { type: 'oceanbase', class: 'oceanbase', text: 'oceanbase' },
  { type: 'pslb', class: 'p_slb', text: '物理负载均衡' },
]

// 其它节点列表
export const otherNodes: Node[] = [
  { type: 'message', class: 'message', text: '短信' },
  { type: 'thirdService', class: 'third_service', text: '其它三方服务' },
  { type: 'cdn', class: 'cdn', text: 'CDN' },
  { type: 'telephone', class: 'telephone', text: '电话' },
  { type: 'email', class: 'email', text: '邮箱' },
  { type: 'slider', class: 'slider', text: '滑块' },
  { type: 'tianyancha', class: 'tianyancha', text: '天眼查' },
  { type: 'deepseek', class: 'deepseek', text: 'DeepSeek' },
  { type: 'doubao', class: 'doubao', text: '豆包' },
]

export const shortStyles = [
  {
    backgroundColor: 'rgb(255, 255, 255)',
    borderWidth: '1px',
    borderColor: 'rgb(42, 42, 42)',
  },
  {
    backgroundColor: 'rgb(245, 245, 245)',
    borderWidth: '1px',
    borderColor: 'rgb(102, 102, 102)',
  },
  {
    backgroundColor: 'rgb(218, 232, 252)',
    borderWidth: '1px',
    borderColor: 'rgb(108, 142, 191)',
  },
  {
    backgroundColor: 'rgb(213, 232, 212)',
    borderWidth: '1px',
    borderColor: 'rgb(130, 179, 102)',
  },
  {
    backgroundColor: 'rgb(255, 230, 204)',
    borderWidth: '1px',
    borderColor: 'rgb(215, 155, 0)',
  },
  {
    backgroundColor: 'rgb(255, 242, 204)',
    borderWidth: '1px',
    borderColor: 'rgb(214, 182, 86)',
  },
  {
    backgroundColor: 'rgb(248, 206, 204)',
    borderWidth: '1px',
    borderColor: 'rgb(184, 84, 80)',
  },
  {
    backgroundColor: 'rgb(220, 210, 230)',
    borderWidth: '1px',
    borderColor: 'rgb(150, 115, 166)',
  },
]

export const borderStyles = [
  {
    value: 'solid',
    label: '',
  },
  {
    value: 'dashed',
    label: '',
  },
  {
    value: 'dotted',
    label: '',
  },
]

export const fontFamilies = [
  {
    value: 'Arial',
    label: '',
  },
  {
    value: 'Verdana',
    label: '',
  },
  {
    value: 'Georgia',
    label: '',
  },
  {
    value: 'Times New Roman',
    label: '',
  },
]

export const GridBackgroundImage =
  'url("data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNDAiIGhlaWdodD0iNDAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PHBhdHRlcm4gaWQ9ImdyaWQiIHdpZHRoPSI0MCIgaGVpZ2h0PSI0MCIgcGF0dGVyblVuaXRzPSJ1c2VyU3BhY2VPblVzZSI+PHBhdGggZD0iTSAwIDEwIEwgNDAgMTAgTSAxMCAwIEwgMTAgNDAgTSAwIDIwIEwgNDAgMjAgTSAyMCAwIEwgMjAgNDAgTSAwIDMwIEwgNDAgMzAgTSAzMCAwIEwgMzAgNDAiIGZpbGw9Im5vbmUiIHN0cm9rZT0iI2QwZDBkMCIgb3BhY2l0eT0iMC4yIiBzdHJva2Utd2lkdGg9IjEiLz48cGF0aCBkPSJNIDQwIDAgTCAwIDAgMCA0MCIgZmlsbD0ibm9uZSIgc3Ryb2tlPSIjZDBkMGQwIiBzdHJva2Utd2lkdGg9IjEiLz48L3BhdHRlcm4+PC9kZWZzPjxyZWN0IHdpZHRoPSIxMDAlIiBoZWlnaHQ9IjEwMCUiIGZpbGw9InVybCgjZ3JpZCkiLz48L3N2Zz4=")'
