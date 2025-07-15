# 部署指南

## 概述

本文档介绍了如何在不同环境中部署 CodeChangeTracker 应用程序。

## 部署方式

### 1. Docker 部署

#### 单容器部署

```bash
# 构建镜像
docker build -t codechange-tracker:latest .

# 运行容器
docker run -d \
  --name codechange-tracker \
  -p 8080:8080 \
  -e DATABASE_URL="postgres://user:pass@host:5432/db" \
  -e REDIS_URL="redis://host:6379" \
  -e JWT_SECRET="your-secret-key" \
  codechange-tracker:latest
```

#### Docker Compose 部署

```bash
# 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f app

# 停止服务
docker-compose down
```

### 2. Kubernetes 部署

#### 准备工作

创建命名空间：

```bash
kubectl create namespace codechange-tracker
```

#### ConfigMap 配置

```yaml
# config/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
  namespace: codechange-tracker
data:
  ENVIRONMENT: "production"
  PORT: "8080"
  LOG_LEVEL: "info"
```

#### Secret 配置

```yaml
# config/secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: app-secrets
  namespace: codechange-tracker
type: Opaque
data:
  DATABASE_URL: <base64-encoded-database-url>
  REDIS_URL: <base64-encoded-redis-url>
  JWT_SECRET: <base64-encoded-jwt-secret>
  OPENAI_API_KEY: <base64-encoded-openai-key>
```

#### 应用部署

```yaml
# deployment/app.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: codechange-tracker
  namespace: codechange-tracker
spec:
  replicas: 3
  selector:
    matchLabels:
      app: codechange-tracker
  template:
    metadata:
      labels:
        app: codechange-tracker
    spec:
      containers:
      - name: app
        image: codechange-tracker:latest
        ports:
        - containerPort: 8080
        envFrom:
        - configMapRef:
            name: app-config
        - secretRef:
            name: app-secrets
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
```

#### Service 配置

```yaml
# service/app-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: codechange-tracker-service
  namespace: codechange-tracker
spec:
  selector:
    app: codechange-tracker
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: ClusterIP
```

#### Ingress 配置

```yaml
# ingress/app-ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: codechange-tracker-ingress
  namespace: codechange-tracker
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  tls:
  - hosts:
    - codechange-tracker.yourdomain.com
    secretName: codechange-tracker-tls
  rules:
  - host: codechange-tracker.yourdomain.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: codechange-tracker-service
            port:
              number: 80
```

#### 部署命令

```bash
# 应用所有配置
kubectl apply -f config/
kubectl apply -f deployment/
kubectl apply -f service/
kubectl apply -f ingress/

# 检查部署状态
kubectl get pods -n codechange-tracker
kubectl get services -n codechange-tracker
kubectl get ingress -n codechange-tracker
```

### 3. 云平台部署

#### AWS ECS

```json
{
  "family": "codechange-tracker",
  "networkMode": "awsvpc",
  "requiresCompatibilities": ["FARGATE"],
  "cpu": "512",
  "memory": "1024",
  "executionRoleArn": "arn:aws:iam::account:role/ecsTaskExecutionRole",
  "containerDefinitions": [
    {
      "name": "app",
      "image": "your-registry/codechange-tracker:latest",
      "portMappings": [
        {
          "containerPort": 8080,
          "protocol": "tcp"
        }
      ],
      "environment": [
        {
          "name": "ENVIRONMENT",
          "value": "production"
        }
      ],
      "secrets": [
        {
          "name": "DATABASE_URL",
          "valueFrom": "arn:aws:ssm:region:account:parameter/codechange-tracker/database-url"
        }
      ],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/codechange-tracker",
          "awslogs-region": "us-west-2",
          "awslogs-stream-prefix": "ecs"
        }
      }
    }
  ]
}
```

#### Google Cloud Run

```yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: codechange-tracker
  annotations:
    run.googleapis.com/ingress: all
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/maxScale: "10"
        run.googleapis.com/cloudsql-instances: project:region:instance
    spec:
      containers:
      - image: gcr.io/project/codechange-tracker:latest
        ports:
        - containerPort: 8080
        env:
        - name: ENVIRONMENT
          value: "production"
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: url
        resources:
          limits:
            cpu: "1"
            memory: "512Mi"
```

## 数据库设置

### PostgreSQL

#### Docker 部署

```bash
docker run -d \
  --name postgres \
  -e POSTGRES_DB=codechange_tracker \
  -e POSTGRES_USER=app_user \
  -e POSTGRES_PASSWORD=secure_password \
  -p 5432:5432 \
  -v postgres_data:/var/lib/postgresql/data \
  postgres:15
```

#### 云数据库配置

**AWS RDS:**

```bash
# 创建数据库实例
aws rds create-db-instance \
  --db-instance-identifier codechange-tracker-db \
  --db-instance-class db.t3.micro \
  --engine postgres \
  --master-username app_user \
  --master-user-password secure_password \
  --allocated-storage 20
```

**Google Cloud SQL:**

```bash
# 创建实例
gcloud sql instances create codechange-tracker-db \
  --database-version=POSTGRES_13 \
  --tier=db-f1-micro \
  --region=us-central1

# 创建数据库
gcloud sql databases create codechange_tracker \
  --instance=codechange-tracker-db
```

### Redis

#### Docker 部署

```bash
docker run -d \
  --name redis \
  -p 6379:6379 \
  -v redis_data:/data \
  redis:7-alpine redis-server --appendonly yes
```

#### 云 Redis 配置

**AWS ElastiCache:**

```bash
aws elasticache create-cache-cluster \
  --cache-cluster-id codechange-tracker-redis \
  --cache-node-type cache.t3.micro \
  --engine redis \
  --num-cache-nodes 1
```

## 监控和日志

### Prometheus 监控

```yaml
# monitoring/prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'codechange-tracker'
    static_configs:
      - targets: ['app:8080']
    metrics_path: /metrics
```

### Grafana 仪表盘

```bash
# 导入预配置的仪表盘
curl -X POST \
  http://grafana:3000/api/dashboards/db \
  -H 'Content-Type: application/json' \
  -d @grafana-dashboard.json
```

### 日志聚合 (ELK Stack)

```yaml
# logging/elasticsearch.yml
version: '3'
services:
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:8.0.0
    environment:
      - discovery.type=single-node
      - xpack.security.enabled=false
    
  logstash:
    image: docker.elastic.co/logstash/logstash:8.0.0
    volumes:
      - ./logstash.conf:/usr/share/logstash/pipeline/logstash.conf
    
  kibana:
    image: docker.elastic.co/kibana/kibana:8.0.0
    ports:
      - "5601:5601"
```

## 安全配置

### HTTPS/TLS

#### Let's Encrypt (Certbot)

```bash
# 安装证书
certbot --nginx -d codechange-tracker.yourdomain.com

# 自动续期
echo "0 12 * * * /usr/bin/certbot renew --quiet" | crontab -
```

#### 自签名证书 (开发环境)

```bash
# 生成证书
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout server.key -out server.crt \
  -subj "/C=US/ST=State/L=City/O=Org/CN=localhost"
```

### 网络安全

#### 防火墙配置

```bash
# UFW 配置
ufw allow 22/tcp    # SSH
ufw allow 80/tcp    # HTTP
ufw allow 443/tcp   # HTTPS
ufw enable
```

#### API 速率限制

在Nginx中配置：

```nginx
http {
    limit_req_zone $binary_remote_addr zone=api:10m rate=10r/s;
    
    server {
        location /api/ {
            limit_req zone=api burst=20 nodelay;
            proxy_pass http://backend;
        }
    }
}
```

## 性能优化

### 负载均衡

#### Nginx 配置

```nginx
upstream backend {
    server app1:8080 weight=3;
    server app2:8080 weight=2;
    server app3:8080 weight=1;
}

server {
    listen 80;
    server_name codechange-tracker.yourdomain.com;
    
    location / {
        proxy_pass http://backend;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### 缓存策略

#### Redis 集群

```yaml
# redis-cluster.yml
version: '3'
services:
  redis-master:
    image: redis:7-alpine
    command: redis-server --appendonly yes
    
  redis-slave:
    image: redis:7-alpine
    command: redis-server --slaveof redis-master 6379
```

## 备份和恢复

### 数据库备份

```bash
#!/bin/bash
# backup.sh

BACKUP_DIR="/backups"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

# PostgreSQL 备份
pg_dump $DATABASE_URL > $BACKUP_DIR/db_backup_$TIMESTAMP.sql

# 压缩备份文件
gzip $BACKUP_DIR/db_backup_$TIMESTAMP.sql

# 清理旧备份 (保留7天)
find $BACKUP_DIR -name "*.sql.gz" -mtime +7 -delete
```

### 自动备份 (Cron)

```bash
# 每天凌晨2点备份
0 2 * * * /opt/scripts/backup.sh
```

## 故障排除

### 常见问题

1. **数据库连接失败**
   ```bash
   # 检查数据库连接
   pg_isready -h $DB_HOST -p $DB_PORT
   ```

2. **Redis 连接问题**
   ```bash
   # 测试 Redis 连接
   redis-cli -h $REDIS_HOST ping
   ```

3. **应用启动失败**
   ```bash
   # 查看容器日志
   docker logs codechange-tracker
   
   # 查看 Kubernetes Pod 日志
   kubectl logs -f deployment/codechange-tracker
   ```

### 健康检查

```bash
# 应用健康检查
curl -f http://localhost:8080/health || exit 1

# 数据库健康检查
pg_isready -h $DB_HOST -p $DB_PORT || exit 1

# Redis 健康检查
redis-cli -h $REDIS_HOST ping || exit 1
```

## 滚动更新

### Docker Swarm

```bash
# 更新服务
docker service update \
  --image codechange-tracker:v2.0.0 \
  codechange-tracker_app
```

### Kubernetes

```bash
# 滚动更新
kubectl set image deployment/codechange-tracker \
  app=codechange-tracker:v2.0.0

# 监控更新进度
kubectl rollout status deployment/codechange-tracker

# 回滚 (如果需要)
kubectl rollout undo deployment/codechange-tracker
```