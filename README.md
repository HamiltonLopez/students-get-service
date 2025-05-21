# Students Get Service

Servicio responsable de listar todos los estudiantes registrados en el sistema.

## Funcionalidad

Este servicio expone un endpoint GET que permite obtener la lista completa de estudiantes almacenados en la base de datos MongoDB.

## Especificaciones Técnicas

- **Puerto**: 8082 (interno), 30082 (NodePort)
- **Endpoint**: GET `/students`
- **Runtime**: Go
- **Base de Datos**: MongoDB

## Estructura del Servicio

```
students-get-service/
├── k8s/
│   ├── deployment.yaml
│   └── service.yaml
├── src/
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   └── config/
├── Dockerfile
└── README.md
```

## API Endpoint

### GET /students

Retorna la lista de todos los estudiantes.

#### Query Parameters
- `limit` (opcional): Número máximo de resultados
- `skip` (opcional): Número de resultados a saltar
- `sort` (opcional): Campo por el cual ordenar

#### Response
```json
{
    "students": [
        {
            "id": "string",
            "name": "string",
            "age": number,
            "email": "string",
            "created_at": "timestamp"
        }
    ],
    "total": number
}
```

## Configuración Kubernetes

### Deployment
- **Replicas**: 3
- **Imagen**: hamiltonlg/students-get-service:latest
- **Variables de Entorno**:
  - MONGO_URI: mongodb://mongo-service:27017

### Service
- **Tipo**: NodePort
- **Puerto**: 8082 -> 30082

## Despliegue

```bash
kubectl apply -f k8s/
```

## Verificación

1. Verificar el deployment:
```bash
kubectl get deployment students-get-deployment
```

2. Verificar los pods:
```bash
kubectl get pods -l app=students-get
```

3. Verificar el servicio:
```bash
kubectl get svc students-get-service
```

## Pruebas

### Obtener todos los estudiantes
```bash
curl http://localhost:30082/students
```

### Obtener estudiantes con paginación
```bash
curl http://localhost:30082/students?limit=10&skip=0
```

## Logs

Ver logs de un pod específico:
```bash
kubectl logs -f <pod-name>
```

## Monitoreo

### Métricas Importantes
- Tiempo de respuesta del endpoint
- Número de registros retornados
- Uso de recursos (CPU/Memoria)
- Latencia de consulta a MongoDB

## Solución de Problemas

1. **Error de Conexión a MongoDB**:
   - Verificar la variable MONGO_URI
   - Comprobar conectividad con mongo-service
   - Revisar logs de MongoDB

2. **Rendimiento Lento**:
   - Verificar índices en MongoDB
   - Comprobar el tamaño de la respuesta
   - Analizar la consulta MongoDB

3. **Pod en CrashLoopBackOff**:
   - Verificar logs del pod
   - Comprobar recursos asignados
   - Verificar configuración del deployment

4. **Servicio no accesible**:
   - Verificar el estado del service
   - Comprobar la configuración de NodePort
   - Verificar reglas de firewall

## Optimización

1. **Índices MongoDB**:
   - Crear índices para campos frecuentemente consultados
   - Mantener estadísticas actualizadas

2. **Caché**:
   - Implementar caché para consultas frecuentes
   - Configurar tiempo de expiración apropiado

3. **Paginación**:
   - Utilizar skip/limit para grandes conjuntos de datos
   - Implementar cursor-based pagination para mejor rendimiento 