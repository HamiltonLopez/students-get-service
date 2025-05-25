# Students Get Service

Este servicio es parte del sistema de gestión de estudiantes y se encarga de listar todos los estudiantes con capacidades de paginación y filtrado.

## Estructura del Servicio

```
students-get-service/
├── controllers/     # Controladores REST
├── models/         # Modelos de datos
├── repositories/   # Capa de acceso a datos
├── services/      # Lógica de negocio
├── k8s/           # Configuraciones de Kubernetes
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
└── test/          # Scripts de prueba
    └── test-get.sh
```

## Endpoints

### GET /students
Lista todos los estudiantes con soporte para paginación.

**Parámetros de Query:**
- `limit`: Número de registros por página (default: 10)
- `skip`: Número de registros a saltar (default: 0)
- `sort`: Campo por el cual ordenar (default: "name")

**Response (200 OK):**
```json
{
    "students": [
        {
            "id": "string",
            "name": "string",
            "age": number,
            "email": "string"
        }
    ],
    "total": number,
    "page": number,
    "pages": number
}
```

## Configuración Kubernetes

### Deployment
El servicio se despliega con las siguientes especificaciones:
- Replicas: 1
- Puerto: 8080
- Imagen: students-get-service:latest

### Service
- Tipo: NodePort
- Puerto: 8080
- NodePort: 30082

### Ingress
- Path: /students
- Servicio: students-get-service
- Puerto: 8080

## Despliegue en Kubernetes

### 1. Aplicar configuraciones
```bash
# Crear el deployment
kubectl apply -f k8s/deployment.yaml

# Crear el service
kubectl apply -f k8s/service.yaml

# Crear el ingress
kubectl apply -f k8s/ingress.yaml
```

### 2. Verificar el despliegue
```bash
# Verificar el deployment
kubectl get deployment students-get-deployment
kubectl describe deployment students-get-deployment

# Verificar los pods
kubectl get pods -l app=students-get
kubectl describe pod -l app=students-get

# Verificar el service
kubectl get svc students-get-service
kubectl describe svc students-get-service

# Verificar el ingress
kubectl get ingress students-get-ingress
kubectl describe ingress students-get-ingress
```

### 3. Verificar logs
```bash
# Ver logs de los pods
kubectl logs -l app=students-get
```

### 4. Escalar el servicio
```bash
# Escalar a más réplicas si es necesario
kubectl scale deployment students-get-deployment --replicas=3
```

### 5. Actualizar el servicio
```bash
# Actualizar la imagen del servicio
kubectl set image deployment/students-get-deployment students-get=students-get-service:nueva-version
```

### 6. Eliminar recursos
```bash
# Si necesitas eliminar los recursos
kubectl delete -f k8s/ingress.yaml
kubectl delete -f k8s/service.yaml
kubectl delete -f k8s/deployment.yaml
```

## Pruebas

El servicio incluye un script de pruebas automatizadas (`test/test-get.sh`) que verifica:

1. Obtención de lista de estudiantes
2. Paginación correcta
3. Ordenamiento de resultados
4. Manejo de parámetros inválidos

Para ejecutar las pruebas:
```bash
./test/test-get.sh
```

También se puede ejecutar como parte de la suite completa de pruebas:
```bash
./test-all-services.sh
```

### Casos de Prueba

1. **Test 1:** Obtener todos los estudiantes
   - Verifica la estructura de la respuesta
   - Comprueba el conteo total de estudiantes

2. **Test 2:** Probar paginación
   - Verifica límite de registros por página
   - Comprueba el funcionamiento del skip

3. **Test 3:** Probar ordenamiento
   - Verifica el ordenamiento por diferentes campos
   - Comprueba el orden ascendente/descendente

4. **Test 4:** Parámetros inválidos
   - Prueba con valores no válidos
   - Verifica el manejo de errores

## Variables de Entorno

- `MONGODB_URI`: URI de conexión a MongoDB (default: "mongodb://mongo-service:27017")
- `DATABASE_NAME`: Nombre de la base de datos (default: "studentsdb")
- `COLLECTION_NAME`: Nombre de la colección (default: "students")
- `PAGE_SIZE`: Tamaño de página por defecto (default: 10)

## Dependencias

- Go 1.19+
- MongoDB
- Kubernetes 1.19+
- Ingress NGINX Controller

## Consideraciones de Seguridad

1. Validación de parámetros de consulta
2. Límites en el tamaño de página
3. Sanitización de parámetros de ordenamiento
4. Manejo seguro de errores

## Monitoreo y Logs

- Endpoint de health check: `/health`
- Logs en formato JSON
- Métricas de rendimiento:
  - Tiempo de respuesta
  - Número de registros procesados
  - Uso de memoria en paginación

## Solución de Problemas

1. Verificar la conexión con MongoDB
2. Comprobar los logs del pod
3. Validar la configuración del Ingress
4. Verificar el estado del servicio en Kubernetes
5. Revisar la configuración de paginación 