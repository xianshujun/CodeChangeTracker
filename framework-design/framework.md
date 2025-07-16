code-risk-analysis/  
├── gateway/                  # 转发层（网关）  
│   ├── middleware/            # 中间件：鉴权、限流、日志等  
│   ├── router/                # 路由配置和转发逻辑  
│   └── main.go                # 网关启动入口  
├── cmd/                      # 服务入口  
│   ├── api/                  # gin API服务  
│   │   └── main.go           # 服务启动入口  
│   ├── worker/               # 异步分析worker  
│   │   └── main.go  
├── internal/                 # 内部模块  
│   ├── analyzer/             # 分析核心模块  
│   │   ├── structure/        # 结构分析  
│   │   │   ├── ast_parser.go # go/ast解析实现  
│   │   │   ├── feature_extractor.go # 特征提取算法  
│   │   │   └── model.go      # 结构特征数据模型  
│   │   ├── semantic/         # 语义风险分析  
│   │   │   ├── llm_client.go # LLM调用客户端  
│   │   │   ├── prompt.go     # Prompt工程  
│   │   │   └── risk_model.go # 业务风险评估模型  
│   │   ├── agent/  
│   │   │   ├── diff.go      # ai过滤无关变更  
│   │   │   └── other.go     #其余设计ai的代码  
│   ├── llm/              # LLM模型  
│   │   └── historical/       # 历史关联分析  
│   │       ├── graph_rag.go  # GraphRAG实现  
│   │       └── kg_store.go   # 知识图谱存储逻辑  
│   ├── preprocessor/         # 代码预处理  
│   │   ├── cutter.go         # 代码切割工具  
│   │   ├── summarizer.go     # 摘要提取  
│   │   └── pruner.go         # 剪枝优化  
│   ├── collector/            # 数据采集  
│   │   ├── git_handler.go    # git2go封装  
│   │   └── data_converter.go # 数据格式转换  
│   ├── storage/              # 存储层  
│   │   ├── vector_db.go      # 向量数据库操作  
│   │   ├── mysql_db.go       # 关系数据库操作  
│   │   └── graph_db.go      # 图数据库操作  
│   ├── service/              # 业务服务层  
│   │   └── api_service.go    # API接口服务  
│   └── config/               # 配置管理  
│       ├── etcd_client.go    # etcd客户端  
│       └── config.go         # 配置模型  
├── pkg/                      # 公共包  
│   ├── logger/               # 日志工具  
│   ├── utils/                # 通用工具函数  
│   └── model/                # 公共数据模型  
├── api/                      # API定义  
│   ├── swagger/              # 接口文档  
│   └── handler/              # 接口处理器  
├── configs/                  # 配置文件模板  
│   ├── app.yaml  
│   └── etcd.yaml  
├── scripts/                  # 部署脚本  
│   ├── k8s/                  # k8s部署配置  
│   └── build.sh              # 构建脚本  
├── go.mod  
└── README.md