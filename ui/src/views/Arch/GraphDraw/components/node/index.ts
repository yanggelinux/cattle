// 基础图形
import CircleNode from './basic/CircleNode.ts'
import RectNode from './basic/RectNode.ts'
import RectRadiusNode from './basic/RectRadiusNode.ts'
import EllipseNode from './basic/EllipseNode.ts'
import TextNode from './basic/TextNode.ts'
import DiamondNode from './basic/DiamondNode.ts'
// path绘制的个性化图形
import CylindeNode from './path/CylindeNode.ts'
import TriangleNode from './path/TriangleNode.ts'
import ParallelogramNode from './path/ParallelogramNode.ts'
import ActorNode from './path/ActorNode.ts'
import StarNode from './path/Star.ts'
import PentagonNode from './path/PentagonNode.ts'
import HexagonNode from './path/HexagonNode.ts'
import SeptagonNode from './path/SeptagonNode.ts'
import HeptagonNode from './path/HeptagonNode.ts'
import TrapezoidNode from './path/TrapezoidNode.ts'
import CrossNode from './path/CrossNode.ts'
import MinusNode from './path/MinusNode.ts'
import TimesNode from './path/TimesNode.ts'
import DivideNode from './path/DivideNode.ts'
// 多边形绘制的箭头
import LeftArrow from './arrow/LeftArrow.ts'
import RightArrow from './arrow/RightArrow.ts'
import HorizontalArrow from './arrow/HorizontalArrowNode.ts'
import UpArrow from './arrow/UpArrowNode.ts'
import DownArrow from './arrow/DownArrowNode.ts'
import VerticalArrow from './arrow/VerticalArrowNode.ts'
// image绘制图片节点
import ImageSetting from './image/Setting.ts'
import ImageUser from './image/User.ts'
import ImageCloud from './image/Cloud.ts'
// image绘制左上角icon节点
import ImageMessage from './image/Message.ts'
import ImageThirdService from './image/ThirdService.ts'
import ImageSlider from './image/Slider.ts'
import ImageTelephone from './image/Telephone.ts'
import ImageTianyancha from './image/Tianyancha.ts'
import ImageEmail from './image/Email.ts'
import ImageBrowser from './image/Browser.ts'
import ImageCDN from './image/CDN.ts'
import ImageMSHA from './image/MSHA.ts'
import ImageEs from './image/ES.ts'
import ImageDeepseek from './image/Deepseek.ts'
import ImageDoubao from './image/Doubao.ts'
import ImageECS from './image/ECS.ts'
import ImageCSB from './image/CSB.ts'
import ImageDomain from './image/Domain.ts'
import ImageDRDS from './image/DRDS.ts'
import ImageGBase from './image/GBase.ts'
import ImageHBase from './image/HBase.ts'
import ImageKong from './image/Kong.ts'
import ImageMQ from './image/MQ.ts'
import ImageJavaApp from './image/JavaApp.ts'
import ImageNginx from './image/Nginx.ts'
import ImageOceanBase from './image/OceanBase.ts'
import ImageUnderCloud from './image/UnderCloud.ts'
import ImageOracle from './image/Oracle.ts'
import ImageOSS from './image/OSS.ts'
import ImageWAF from './image/WAF.ts'
import ImageRDS from './image/RDS.ts'
import ImageRedis from './image/Redis.ts'
import ImageSLB from './image/SLB.ts'
import ImageSSLVPN from './image/SSLVPN.ts'
import ImageSwitch from './image/Switch.ts'
import ImageFirewall from './image/Firewall.ts'
import ImageInternalFace from './image/InternalFace.ts'
import ImageEncryptMachine from './image/EncryptMachine.ts'
import ImagePhysicalMachine from './image/PhysicalMachine.ts'
import ImageIntegratedMachine from './image/IntegratedMachine.ts'
import ImageRouter from './image/Router.ts'
import ImageLocal from './image/Local.ts'
import ImageArbitration from './image/Arbitration'
import ImageMpass from './image/Mpass.ts'
import ImageAdapter from './image/Adapter.ts'
import ImageArchGraph from './image/ArchGraph.ts'
import ImageSls from './image/Sls.ts'
import ImageMinistry from './image/Ministry.ts'
import ImageTencentCloud from './image/TencentCloud.ts'
import ImagePSLB from './image/PSLB.ts'

// group

import RectGroup from './group/RectGroup.ts'
import CircleGroup from './group/CircleGroup.ts'
import EllipseGroup from './group/EllipseGroup.ts'
import TriangleGroup from './group/TriangleGroup.ts'

// network

import NetworkGroup from './network/NetWorkGroup.ts'

// 注册边
import Ployline from './edge/Polyline.ts'
import Line from './edge/Line.ts'
import Bezier from './edge/Bezier.ts'

// proc流程节点

import ProcStart from './path/Start.ts'
import ProcEnd from './path/End.ts'
import ProcApproval from './path/Approval.ts'

export const registerCustomElement = (lf: any) => {
  // 注册基础图形
  lf.register(CircleNode)
  lf.register(RectNode)
  // lf.register(CustomRectNode)
  lf.register(RectRadiusNode)
  lf.register(EllipseNode)
  lf.register(DiamondNode)
  lf.register(TextNode)
  // 注册path绘制的个性化图形
  lf.register(CylindeNode)
  lf.register(TriangleNode)
  lf.register(ParallelogramNode)
  lf.register(ActorNode)
  lf.register(StarNode)
  lf.register(PentagonNode)
  lf.register(HexagonNode)
  lf.register(SeptagonNode)
  lf.register(HeptagonNode)
  lf.register(TrapezoidNode)
  lf.register(CrossNode)
  lf.register(MinusNode)
  lf.register(TimesNode)
  lf.register(DivideNode)
  // 注册多边形绘制的箭头
  lf.register(LeftArrow)
  lf.register(RightArrow)
  lf.register(HorizontalArrow)
  lf.register(UpArrow)
  lf.register(DownArrow)
  lf.register(VerticalArrow)
  // 注册image绘制图片节点
  lf.register(ImageSetting)
  lf.register(ImageUser)
  lf.register(ImageCloud)

  // 注册image绘制左上角icon节点
  lf.register(ImageMessage)
  lf.register(ImageThirdService)
  lf.register(ImageSlider)
  lf.register(ImageTelephone)
  lf.register(ImageTianyancha)
  lf.register(ImageEmail)
  lf.register(ImageCDN)
  lf.register(ImageBrowser)
  lf.register(ImageEs)
  lf.register(ImageDeepseek)
  lf.register(ImageDoubao)
  lf.register(ImageMSHA)
  lf.register(ImageECS)
  lf.register(ImageCSB)
  lf.register(ImageDRDS)
  lf.register(ImageGBase)
  lf.register(ImageKong)
  lf.register(ImageDomain)
  lf.register(ImageMQ)
  lf.register(ImageWAF)
  lf.register(ImageNginx)
  lf.register(ImageOceanBase)
  lf.register(ImageOracle)
  lf.register(ImageOSS)
  lf.register(ImageRDS)
  lf.register(ImageRedis)
  lf.register(ImageSLB)
  lf.register(ImageUnderCloud)
  lf.register(ImageSSLVPN)
  lf.register(ImageSwitch)
  lf.register(ImageInternalFace)
  lf.register(ImageEncryptMachine)
  lf.register(ImagePhysicalMachine)
  lf.register(ImageIntegratedMachine)
  lf.register(ImageFirewall)
  lf.register(ImageJavaApp)
  lf.register(ImageRouter)
  lf.register(ImageLocal)
  lf.register(ImageHBase)
  lf.register(ImageArbitration)
  lf.register(ImageMpass)
  lf.register(ImageAdapter)
  lf.register(ImageArchGraph)
  lf.register(ImageSls)
  lf.register(ImageMinistry)
  lf.register(ImageTencentCloud)
  lf.register(ImagePSLB)

  //group
  lf.register(RectGroup)
  lf.register(CircleGroup)
  lf.register(EllipseGroup)
  lf.register(TriangleGroup)

  //network
  lf.register(NetworkGroup)

  // 注册边
  lf.register(Ployline)
  lf.register(Line)
  lf.register(Bezier)

  // process
  lf.register(ProcStart)
  lf.register(ProcEnd)
  lf.register(ProcApproval)
}
