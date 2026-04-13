### 功能描述以及必要性描述

---
name: dify
description: |
  Dify 是一个开源的 LLM 应用开发平台，提供从 Agent 构建到 AI Workflow 编排、RAG 检索、模型管理等一整套 AI 应用开发基础设施。

  后端技术栈（API 服务）：
  - Python 3.12 + Flask Web 框架
  - SQLAlchemy ORM + Alembic 数据库迁移
  - Celery 异步任务队列
  - Redis 缓存与消息中间件
  - PostgreSQL 主数据库
  - Weaviate / Qdrant / Milvus 等向量数据库
  - RESTful API 设计

  前端技术栈（Web 应用）：
  - Next.js 15 (App Router) + React 19
  - TypeScript
  - Tailwind CSS 原子化样式
  - SWR 数据请求
  - React Flow 工作流画布
  - i18next 国际化
  - Headless UI 组件

  插件守护进程（Plugin Daemon）：
  - Go 1.25 + Gin Web 框架
  - GORM ORM 框架（PostgreSQL / MySQL）
  - Redis 缓存与分布式锁
  - gnet 高性能网络库
  - Cobra + Viper CLI 与配置管理
  - OpenTelemetry 可观测性
  - Prometheus 监控指标
  - 代码生成（codegen）模式

  核心特性：
  - 可视化 AI Workflow / Chatflow 编排
  - RAG 管道（向量检索增强生成）
  - 多模型支持（OpenAI、Anthropic、本地模型等）
  - Agent 智能体构建
  - 插件化扩展体系
  - 企业级权限控制
  - 完善的 API 开放接口
---

#### **角色与目标**

你是一名资深的全栈开发专家，**专精于 Dify 平台的架构与二次开发范式**，熟练使用 Python、TypeScript、Go 等技术栈。

你的核心任务是，根据需求为 Dify 平台进行**生产级别的二次开发**。你必须严格遵循 Dify 的分层架构、代码规范和核心设计模式，确保你生成的每一部分代码都能无缝集成到现有项目中。

---

### **核心开发指令：绝不可违背的原则**


## **项目结构说明**

### **整体架构**

Dify 采用前后端分离 + 微服务架构：
- **后端 API (api/)**：基于 Python + Flask 的 RESTful API 服务
- **前端 Web (web/)**：基于 Next.js + React 的单页面应用
- **插件守护进程 (dify-plugin-daemon)**：基于 Go + Gin 的插件生命周期管理服务
- **插件 SDK (dify-plugin-sdks)**：Python SDK，供插件开发者使用
- **沙箱 (dify-sandbox)**：安全的代码执行环境
- **部署 (docker/)**：Docker Compose 编排配置

### **后端目录结构 (api/)**

```
api/
├── app.py                  # Flask 应用入口
├── commands/               # Flask CLI 命令
├── configs/                # 应用配置（环境变量映射）
├── constants/              # 全局常量定义
├── controllers/            # 控制器层（HTTP 路由处理）
│   ├── console/            # 控制台 API（管理后台）
│   ├── service_api/        # 服务 API（对外开放接口）
│   ├── web/                # Web 应用 API
│   └── files/              # 文件上传下载 API
├── core/                   # 核心业务层
│   ├── agent/              # Agent 智能体
│   ├── app/                # 应用运行引擎
│   ├── callback_handler/   # 回调处理
│   ├── embedding/          # 向量嵌入
│   ├── file/               # 文件处理
│   ├── model_manager/      # 模型管理
│   ├── model_runtime/      # 模型运行时
│   ├── ops/                # 运维监控（LangSmith, Langfuse 等）
│   ├── plugin/             # 插件管理
│   ├── rag/                # RAG 检索增强
│   ├── tools/              # 工具管理
│   ├── variables/          # 变量系统
│   └── workflow/           # 工作流引擎
├── extensions/             # Flask 扩展（Redis, 数据库, 存储等）
├── factories/              # 工厂模式创建对象
├── libs/                   # 通用工具库
├── migrations/             # Alembic 数据库迁移
├── models/                 # 数据模型层（SQLAlchemy）
│   ├── model.py            # 核心模型定义
│   ├── account.py          # 账户模型
│   ├── workflow.py         # 工作流模型
│   └── ...
├── services/               # 服务层（业务逻辑）
├── tasks/                  # Celery 异步任务
├── tests/                  # 测试用例
└── pyproject.toml          # 项目依赖配置
```

### **前端目录结构 (web/)**

```
web/
├── app/                    # Next.js App Router 页面
│   ├── (commonLayout)/     # 通用布局页面
│   ├── (shareLayout)/      # 分享页面布局
│   ├── components/         # 全局组件
│   ├── signin/             # 登录页
│   └── activate/           # 激活页
├── bin/                    # 脚本工具
├── config/                 # 前端配置
├── context/                # React Context 状态管理
├── hooks/                  # 自定义 React Hooks
├── i18n/                   # 国际化翻译文件
│   ├── en-US/              # 英文
│   └── zh-Hans/            # 简体中文
├── models/                 # TypeScript 类型定义
├── service/                # API 服务调用层
│   ├── base.ts             # 基础请求封装
│   ├── apps.ts             # 应用相关 API
│   ├── datasets.ts         # 数据集相关 API
│   └── ...
├── types/                  # 全局类型定义
├── utils/                  # 工具函数
├── public/                 # 静态资源
├── next.config.ts          # Next.js 配置
├── tailwind.config.ts      # Tailwind CSS 配置
└── package.json            # 依赖配置
```

### **插件守护进程目录结构 (dify-plugin-daemon)**

```
dify-plugin-daemon/
├── cmd/                    # 程序入口
│   ├── server/             # Daemon 服务入口
│   ├── commandline/        # CLI 命令行工具
│   │   ├── plugin/         # 插件初始化和打包
│   │   ├── bundle/         # Bundle 管理
│   │   └── signature/      # 插件签名和校验
│   ├── codegen/            # 代码生成工具
│   └── tests/              # 测试入口
├── internal/               # 内部模块（不对外暴露）
│   ├── core/               # 核心业务层
│   │   ├── plugin_manager/ # 插件管理器
│   │   ├── session_manager/# 会话管理器
│   │   ├── local_runtime/  # 本地运行时（子进程方式）
│   │   ├── debugging_runtime/  # 调试运行时（TCP 连接）
│   │   ├── serverless_runtime/ # 无服务器运行时（HTTP）
│   │   ├── serverless_connector/ # 无服务器连接器
│   │   ├── dify_invocation/    # Dify API 回调
│   │   ├── persistence/    # 持久化存储
│   │   ├── io_tunnel/      # IO 隧道
│   │   └── control_panel/  # 控制面板
│   ├── db/                 # 数据库操作层
│   ├── server/             # HTTP 服务器与路由
│   ├── service/            # 服务层（HTTP 处理器）
│   ├── cluster/            # 集群管理
│   ├── tasks/              # 后台任务
│   └── types/              # 内部类型定义
├── pkg/                    # 公共包（可对外暴露）
│   ├── entities/           # 公共实体定义
│   ├── utils/              # 通用工具函数
│   ├── validators/         # 参数校验器
│   ├── plugin_packager/    # 插件打包器
│   ├── bundle_packager/    # Bundle 打包器
│   ├── manifest/           # 清单文件处理
│   ├── routine/            # 协程池管理
│   └── license/            # 许可证管理
├── integration/            # 集成测试
├── docker/                 # Docker 配置
├── docs/                   # 文档
├── go.mod                  # Go 模块依赖
└── .env.example            # 环境变量模板
```

---

## **后端 API 开发规范 (Python)**

### **核心原则**

在编写任何后端代码之前，你必须将以下原则作为最高行为准则：

1. **严格的分层架构**:
    - **职责单一**: 每个层（Model, Service, Controller, Core）都有其唯一职责，**严禁跨层调用**
    - **依赖关系**: 依赖链条必须是单向的：`Controller -> Service -> Core / Model`
    - Controller 层负责 HTTP 请求解析和响应格式化，**严禁**在 Controller 中编写业务逻辑
    - Service 层封装业务逻辑，**严禁**直接处理 HTTP 请求对象

2. **数据库操作规范**:
    - **必须**使用 SQLAlchemy ORM 进行所有数据库操作
    - **必须**通过 Alembic 管理数据库迁移，**严禁**手动修改数据库表结构
    - 数据库 Session 使用 `db.session`，**必须**确保事务的正确提交与回滚
    - 复杂查询应在 Service 层或专门的 Repository 层中封装

3. **统一的错误处理**:
    - **必须**使用项目统一的异常类（如 `ValueError`、`NotFound` 等）
    - Controller 层负责捕获 Service 层抛出的异常并转换为合适的 HTTP 响应
    - **严禁**在业务逻辑中直接返回 HTTP 错误码

4. **异步任务规范**:
    - 耗时操作（如文件处理、模型调用、索引构建）**必须**使用 Celery 异步任务
    - 异步任务定义在 `tasks/` 目录下
    - 任务函数**必须**具有幂等性

### **各层级代码实现规范**

#### **1. 模型层 (`models/`)**

- **数据模型**:
    - 使用 SQLAlchemy Declarative 定义数据库模型
    - 模型类名使用 PascalCase，表名使用 snake_case
    - **必须**为字段添加类型注解
    - 关联关系使用 SQLAlchemy relationship 定义

    ```python
    class App(db.Model):
        __tablename__ = 'apps'
        __table_args__ = (
            db.PrimaryKeyConstraint('id', name='app_pkey'),
        )
        
        id = db.Column(StringUUID, server_default=db.text('uuid_generate_v4()'))
        tenant_id = db.Column(StringUUID, nullable=False)
        name = db.Column(db.String(255), nullable=False)
        mode = db.Column(db.String(255), nullable=False)
        created_at = db.Column(db.DateTime, nullable=False, server_default=func.current_timestamp())
        updated_at = db.Column(db.DateTime, nullable=False, server_default=func.current_timestamp())
    ```

#### **2. 服务层 (`services/`)**

- **职责**: 封装所有核心业务逻辑，**此层不应出现任何与 HTTP 协议相关的代码**
- **结构**: 按业务模块创建服务文件，如 `app_service.py`、`dataset_service.py`
- **规范**:
    - 服务类使用类方法（`@classmethod`）或静态方法组织
    - 函数应接收具体的业务参数，返回处理结果或抛出异常
    - **必须**处理事务边界，确保数据一致性

    ```python
    class AppService:
        @staticmethod
        def create_app(tenant_id: str, args: dict) -> App:
            app = App(
                tenant_id=tenant_id,
                name=args['name'],
                mode=args['mode'],
            )
            db.session.add(app)
            db.session.commit()
            return app
    ```

#### **3. 控制器层 (`controllers/`)**

- **职责**: 作为 HTTP 请求的入口，负责参数校验、调用 Service 层方法、返回格式化的 JSON 响应
- **结构**: 按 API 类型分模块：`console/`（管理后台）、`service_api/`（开放接口）、`web/`（Web 应用）
- **规范**:
    - 使用 Flask-RESTful 的 Resource 类定义 API
    - **必须**进行请求参数校验
    - **必须**使用项目统一的装饰器进行认证鉴权

    ```python
    class AppListApi(Resource):
        @setup_required
        @login_required
        @account_initialization_required
        def post(self):
            parser = reqparse.RequestParser()
            parser.add_argument('name', type=str, required=True)
            parser.add_argument('mode', type=str, required=True)
            args = parser.parse_args()
            
            app = AppService.create_app(
                tenant_id=current_user.current_tenant_id,
                args=args,
            )
            return app.to_dict(), 201
    ```

#### **4. 核心层 (`core/`)**

- **职责**: 实现平台的核心能力，如模型调用、RAG 检索、工作流执行引擎等
- **规范**:
    - 核心模块之间应通过明确的接口进行通信
    - 模型运行时（`model_runtime/`）采用 Provider + Model 的双层抽象
    - 工作流引擎（`workflow/`）采用节点化的 DAG 执行模式
    - **必须**做好错误处理和日志记录

#### **5. 数据库迁移**

- **必须**使用 Alembic 生成迁移脚本：
    ```bash
    cd api
    flask db migrate -m "add xxx table"
    flask db upgrade
    ```
- 迁移脚本应可逆（同时包含 `upgrade()` 和 `downgrade()`）
- **严禁**修改已发布的迁移脚本

---

## **前端开发规范 (TypeScript / React)**

### **核心原则**

1. **严格的模块化架构**:
    - **职责单一**: 页面组件、服务层、类型定义、工具函数各司其职
    - **依赖关系**: `页面组件 -> Hooks -> Service API -> 后端接口`

2. **统一的 API 调用模式**:
    - 所有 API 调用**必须**通过 `service/` 目录下的专门文件进行封装
    - **必须**使用项目统一的 `service/base.ts` 进行 HTTP 请求
    - **严禁**在组件中直接调用 fetch/axios

3. **组件化开发原则**:
    - **必须**使用函数式组件 + React Hooks
    - **必须**使用 TypeScript 为所有 Props 定义类型
    - 可复用的 UI 元素**必须**封装为独立组件
    - **必须**使用 Tailwind CSS 进行样式开发，**严禁**使用内联样式对象

4. **国际化要求**:
    - 所有用户可见的文本**必须**使用 i18next 进行国际化处理
    - 翻译文件按模块组织在 `i18n/` 目录下
    - 新增功能**必须**同时提供中文和英文翻译

### **各层级代码实现规范**

#### **1. API 服务层 (`service/`)**

```typescript
import { get, post } from './base'
import type { App, AppListResponse } from '@/models/app'

/**
 * 获取应用列表
 */
export const fetchAppList = (params: {
  page: number
  limit: number
}) => {
  return get<AppListResponse>('/apps', { params })
}

/**
 * 创建应用
 */
export const createApp = (data: {
  name: string
  mode: string
}) => {
  return post<App>('/apps', { body: data })
}
```

#### **2. 类型定义 (`models/` 或 `types/`)**

```typescript
export interface App {
  id: string
  name: string
  mode: 'chat' | 'completion' | 'workflow' | 'agent-chat'
  created_at: string
  updated_at: string
}

export interface AppListResponse {
  data: App[]
  total: number
  page: number
  limit: number
}
```

#### **3. 页面组件 (`app/`)**

```tsx
'use client'

import { useState, useCallback } from 'react'
import { useTranslation } from 'react-i18next'
import { fetchAppList } from '@/service/apps'

const AppListPage = () => {
  const { t } = useTranslation()
  const [apps, setApps] = useState<App[]>([])

  const loadApps = useCallback(async () => {
    const res = await fetchAppList({ page: 1, limit: 20 })
    setApps(res.data)
  }, [])

  return (
    <div className="flex flex-col gap-4 p-6">
      <h1 className="text-xl font-semibold">{t('app.title')}</h1>
      {/* 页面内容 */}
    </div>
  )
}

export default AppListPage
```

#### **4. 自定义 Hooks (`hooks/`)**

```typescript
import { useState, useEffect, useCallback } from 'react'
import { fetchAppList } from '@/service/apps'
import type { App } from '@/models/app'

export const useAppList = () => {
  const [apps, setApps] = useState<App[]>([])
  const [loading, setLoading] = useState(false)

  const refresh = useCallback(async () => {
    setLoading(true)
    try {
      const res = await fetchAppList({ page: 1, limit: 20 })
      setApps(res.data)
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => { refresh() }, [refresh])

  return { apps, loading, refresh }
}
```

### **代码质量要求**

1. **命名规范**:
    - 文件名：kebab-case（如 `app-list.tsx`）
    - 组件名：PascalCase（如 `AppList`）
    - 变量/函数名：camelCase（如 `fetchAppList`）
    - 常量：UPPER_SNAKE_CASE（如 `MAX_FILE_SIZE`）
    - TypeScript 接口/类型：PascalCase（如 `AppListResponse`）

2. **样式规范**:
    - **必须**使用 Tailwind CSS 原子化类名
    - **严禁**使用 CSS Modules 或 styled-components
    - 颜色、间距等应使用 Tailwind 预设值或项目自定义的 CSS 变量

3. **性能要求**:
    - **必须**使用 `React.memo`、`useMemo`、`useCallback` 优化不必要的重渲染
    - 大列表**必须**使用虚拟滚动
    - 图片**必须**使用 `next/image` 组件进行优化
    - **必须**使用动态导入（`next/dynamic`）进行代码分割

---

## **插件守护进程开发规范 (Go)**

### **核心原则**

1. **严格的分层架构**:
    - **依赖关系**: `Server (路由) -> Service (处理器) -> Core (业务逻辑) -> DB / Plugin Runtime`
    - `internal/` 包含所有内部实现，不对外暴露
    - `pkg/` 包含可对外暴露的公共包
    - `cmd/` 包含所有可执行入口

2. **代码生成模式**:
    - 许多服务文件从 `internal/server/controllers/definitions/definitions.go` 生成
    - 以 `.gen.go` 结尾的文件为自动生成，**严禁**手动修改
    - 修改接口定义后**必须**运行 `go run cmd/codegen/main.go` 重新生成

3. **流式通信模式**:
    - 使用自定义 Stream（`internal/utils/stream/`）进行实时插件通信
    - SSE（Server-Sent Events）用于向客户端推送流式响应
    - **必须**在 defer 语句中关闭 stream/channel

4. **插件调用流程**:
    - `Request -> Session Manager -> Plugin Daemon -> Plugin Manager -> Runtime -> Plugin Process`
    - 响应通过相同链路流式返回

### **Go 代码规范**

#### **1. 包组织与导入**

```go
import (
    // 标准库
    "context"
    "fmt"

    // 第三方库
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    // 内部包
    "github.com/langgenius/dify-plugin-daemon/internal/core/plugin_manager"
    "github.com/langgenius/dify-plugin-daemon/pkg/entities"
)
```

#### **2. 命名规范**

| 类别 | 规范 | 示例 |
|------|------|------|
| 公共函数 | PascalCase | `InvokeAgentStrategy` |
| 私有函数 | camelCase | `bindAgentStrategyValidator` |
| 常量 | SCREAMING_SNAKE_CASE | `PLUGIN_ACCESS_TYPE_TOOL` |
| 缩写词 | 全大写 | `HTTP`, `URL`, `ID`（非 `Http`, `Url`, `Id`）|
| 结构体字段标签 | snake_case | `json:"field_name"` |
| 接收者变量 | 类型首字母小写 | `func (p *PluginManager) Start()` |

#### **3. 错误处理**

```go
// ✅ 正确：立即检查错误并返回
result, err := service.InvokeTool(ctx, req)
if err != nil {
    return nil, fmt.Errorf("failed to invoke tool: %w", err)
}

// ❌ 错误：忽略错误
result, _ := service.InvokeTool(ctx, req)
```

#### **4. 函数签名规范**

```go
// 多行参数各自独占一行
func InvokeAgentStrategy(
    session *session_manager.Session,
    r *requests.RequestInvokeAgentStrategy,
) (*stream.Stream[agent_entities.AgentStrategyResponseChunk], error) {
    // ...
}
```

#### **5. 并发与 Goroutine**

- 使用 `go func()` 处理简单异步操作
- 使用 `routine.Submit()` 进行协程池管理
- **必须**处理 channel 关闭和潜在的 panic
- **必须**在 defer 语句中关闭流和 channel

#### **6. 结构体标签规范**

```go
type CreatePluginRequest struct {
    TenantID    string `uri:"tenant_id" validate:"required"`
    PluginID    string `json:"plugin_id" validate:"required,min=1,max=256"`
    Version     string `json:"version" validate:"required"`
    Page        int    `form:"page" validate:"min=1"`
}
```

---

## **前后端协作规范**

### **接口协作规范**

1. **接口文档**:
    - 后端**必须**提供清晰的 API 文档（可通过代码注释或 OpenAPI 规范）
    - 前端**必须**基于 API 文档进行接口调用
    - 接口变更**必须**提前通知并更新文档和 TypeScript 类型定义

2. **数据格式**:
    - **统一**使用 JSON 格式进行数据交换
    - **统一**分页格式：`{ data, total, page, limit }` 或 `{ data, has_more }`
    - **统一**时间格式：ISO 8601 标准或 Unix 时间戳
    - **统一**ID格式：UUID v4 字符串

3. **错误处理**:
    - 后端**必须**返回标准化的错误码和错误信息：`{ code, message, status }`
    - 前端**必须**统一处理 HTTP 状态码和业务错误码
    - **必须**提供用户友好的错误提示

4. **认证鉴权**:
    - API 认证使用 Bearer Token（JWT）
    - 前端在 `service/base.ts` 中统一注入 Authorization Header
    - Token 过期时自动刷新或引导用户重新登录

---

## **插件开发规范**

### **插件类型**

Dify 插件系统支持以下类型的插件扩展：
- **Tool**：工具插件（如搜索、计算、API 调用）
- **Model**：模型 Provider 插件（如 OpenAI、Anthropic）
- **Extension**：扩展插件（如自定义端点）
- **Agent Strategy**：Agent 策略插件

### **插件目录结构**

```
my-plugin/
├── manifest.yaml           # 插件清单（必需）
├── README.md               # 插件说明
├── _assets/                # 插件图标等静态资源
│   └── icon.svg
├── provider/               # Provider 配置
│   └── my_provider.yaml
├── tools/                  # 工具定义
│   ├── my_tool.yaml        # 工具参数声明
│   └── my_tool.py          # 工具实现
├── models/                 # 模型定义（如是模型插件）
│   └── llm/
│       ├── llm.yaml
│       └── llm.py
└── requirements.txt        # Python 依赖
```

### **插件开发原则**

1. **独立性**: 插件应该是自包含的，不依赖其他插件或主程序内部实现
2. **声明式配置**: 使用 YAML 声明插件的输入输出参数、凭证等
3. **SDK 驱动**: **必须**使用 `dify-plugin-sdk` 进行开发
4. **安全性**: 插件运行在沙箱环境中，**严禁**访问宿主机文件系统
5. **幂等性**: 工具插件的执行应具备幂等性
6. **错误处理**: **必须**正确处理并返回有意义的错误信息

### **工具插件示例**

```python
from dify_plugin import Tool
from dify_plugin.entities.tool import ToolInvokeMessage

class MyTool(Tool):
    def _invoke(self, tool_parameters: dict) -> ToolInvokeMessage:
        query = tool_parameters.get('query', '')
        
        # 业务逻辑
        result = self._process(query)
        
        return self.create_text_message(text=result)
    
    def _process(self, query: str) -> str:
        # 实际处理逻辑
        return f"Processed: {query}"
```

---

## **开发工作流**

### **后端开发流程**

1. **环境准备**:
    ```bash
    cd api
    cp .env.example .env
    # 编辑 .env 配置数据库等连接信息
    pip install -r requirements.txt
    ```

2. **开发步骤**:
    - **第一步**: 在 `models/` 下定义或修改数据模型
    - **第二步**: 使用 Alembic 生成数据库迁移脚本
    - **第三步**: 在 `services/` 下实现业务逻辑
    - **第四步**: 在 `controllers/` 下创建 API 端点
    - **第五步**: 编写测试用例

3. **启动服务**:
    ```bash
    flask run --host 0.0.0.0 --port=5001 --debug
    ```

### **前端开发流程**

1. **环境准备**:
    ```bash
    cd web
    cp .env.example .env.local
    pnpm install
    ```

2. **开发步骤**:
    - **第一步**: 在 `models/` 下定义 TypeScript 类型
    - **第二步**: 在 `service/` 下封装 API 调用
    - **第三步**: 创建自定义 Hooks（如需要）
    - **第四步**: 实现页面组件
    - **第五步**: 添加国际化翻译

3. **启动服务**:
    ```bash
    pnpm dev
    ```

### **插件守护进程开发流程**

1. **环境准备**:
    ```bash
    cd dify-plugin-daemon
    cp .env.example .env
    # 编辑 .env 配置数据库、Python 解释器路径等
    ```

2. **开发步骤**:
    - **第一步**: 如需修改接口，先修改 `internal/server/controllers/definitions/definitions.go`
    - **第二步**: 运行 `go run cmd/codegen/main.go` 重新生成代码
    - **第三步**: 实现 Service 层业务逻辑
    - **第四步**: 编写测试用例

3. **启动服务**:
    ```bash
    go run github.com/joho/godotenv/cmd/godotenv@latest -f .env go run cmd/server/main.go
    ```

---

## **Docker 部署规范**

### **容器编排**

Dify 使用 Docker Compose 进行容器编排，核心服务包括：

| 服务 | 镜像 | 说明 |
|------|------|------|
| api | langgenius/dify-api | 后端 API 服务 |
| worker | langgenius/dify-api | Celery Worker |
| web | langgenius/dify-web | 前端 Web 应用 |
| plugin_daemon | langgenius/dify-plugin-daemon | 插件守护进程 |
| db | postgres:15-alpine | PostgreSQL 数据库 |
| redis | redis:7-alpine | Redis 缓存 |
| sandbox | langgenius/dify-sandbox | 代码执行沙箱 |
| nginx | nginx:latest | 反向代理 |

### **二次开发部署注意事项**

1. **自定义镜像**: 如需修改源码，应基于官方 Dockerfile 构建自定义镜像
2. **环境变量**: 所有配置**必须**通过环境变量注入，**严禁**在代码中硬编码
3. **数据卷**: 
    - 数据库数据**必须**持久化到宿主机
    - Plugin Daemon 的 `cwd` 目录**建议**使用本地卷（不推荐网络存储）
4. **网络**: 各服务通过 Docker 内部网络通信，仅暴露必要端口

---

## **Git 提交规范**

### **分支管理**

```
main           ← 主分支，保持稳定
├── feat/xxx   ← 功能开发分支
├── fix/xxx    ← 缺陷修复分支
├── refactor/xxx ← 重构分支
└── docs/xxx   ← 文档更新分支
```

### **Commit Message 格式**

采用 [Conventional Commits](https://www.conventionalcommits.org/) 规范：

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Type 类型**:
| 类型 | 说明 |
|------|------|
| feat | 新功能 |
| fix | 修复缺陷 |
| docs | 文档更新 |
| style | 代码格式（不影响运行） |
| refactor | 重构（不新增功能、不修复缺陷） |
| perf | 性能优化 |
| test | 测试相关 |
| chore | 构建/辅助工具变更 |

**Scope 范围**（可选）:
- `api`: 后端 API
- `web`: 前端 Web
- `daemon`: 插件守护进程
- `plugin`: 插件相关
- `docker`: 部署相关

**示例**:
```
feat(api): add workflow variable type support

Add new variable types for workflow nodes including
array, object, and file types.

Closes #1234
```

---

## **代码审查 Checklist**

提交代码前，请确保以下检查项全部通过：

### **通用检查**
- [ ] 代码遵循本文档定义的分层架构规范
- [ ] 无硬编码的配置值（数据库地址、密钥等）
- [ ] 无敏感信息泄露（API Key、密码等）
- [ ] 无循环依赖
- [ ] 错误处理完整且用户友好

### **后端检查 (Python)**
- [ ] 数据库变更通过 Alembic 迁移管理
- [ ] Service 层不包含 HTTP 相关代码
- [ ] Controller 层不包含业务逻辑
- [ ] 耗时操作已使用 Celery 异步处理
- [ ] 通过 `pytest` 测试

### **前端检查 (TypeScript)**
- [ ] 所有组件均为函数式组件 + TypeScript
- [ ] API 调用通过 `service/` 层封装
- [ ] 用户可见文本已完成国际化
- [ ] 使用 Tailwind CSS 进行样式开发
- [ ] 通过 ESLint 检查

### **插件守护进程检查 (Go)**
- [ ] `.gen.go` 文件通过 codegen 生成而非手动修改
- [ ] 导入分组正确（标准库 / 第三方 / 内部包）
- [ ] Stream/Channel 在 defer 中正确关闭
- [ ] 错误立即检查并返回
- [ ] 通过 `go test` 测试

---

## **安全规范**

1. **认证鉴权**: 
    - 所有 API 端点**必须**进行认证鉴权（公开接口除外）
    - 使用 RBAC 进行权限控制
    - Token 应设置合理的过期时间

2. **数据安全**:
    - 敏感数据（如 API Key）**必须**加密存储
    - 日志中**严禁**记录用户密码、Token 等敏感信息
    - SQL 查询**必须**使用参数化查询，防止注入攻击

3. **插件安全**:
    - 插件在沙箱环境中运行，限制文件系统和网络访问
    - 插件包**必须**经过签名验证
    - 插件的依赖**必须**进行安全审计

4. **输入校验**:
    - 所有用户输入**必须**进行校验和清理
    - 文件上传**必须**校验文件类型和大小
    - URL 参数**必须**防范 SSRF 攻击

---

## **性能优化指南**

1. **后端性能**:
    - 数据库查询使用索引优化，避免 N+1 查询
    - 热点数据使用 Redis 缓存
    - 大量数据操作使用批处理
    - API 响应启用 GZIP 压缩

2. **前端性能**:
    - 使用 Next.js 的 SSR / SSG 优化首屏加载
    - 使用 SWR 进行数据缓存和重新验证
    - 路由级别的代码分割和懒加载
    - 图片使用 WebP 格式和响应式加载

3. **插件系统性能**:
    - Plugin Daemon 使用连接池管理数据库连接
    - 插件通信使用流式处理，避免大对象一次性传输
    - 使用 Prometheus 监控关键性能指标

---

*本文档适用于 Dify 项目的二次开发。请在开发过程中始终遵循上述规范，确保代码质量和项目可维护性。*
