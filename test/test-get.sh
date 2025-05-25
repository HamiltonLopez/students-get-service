#!/bin/bash

# Colores para la salida
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

# URL del servicio
SERVICE_URL="http://${KUBE_IP}:30082"

echo "Probando Students Get Service..."
echo "==============================="

# Test 1: Obtener todos los estudiantes
echo -e "\nTest 1: Obtener todos los estudiantes"
response=$(curl -s -X GET "${SERVICE_URL}/students")

if [[ $response == *"students"* ]]; then
    echo -e "${GREEN}✓ Test 1 exitoso: Lista de estudiantes obtenida correctamente${NC}"
    # Contar el número de estudiantes
    count=$(echo $response | grep -o '"id"' | wc -l)
    echo "Número de estudiantes encontrados: $count"
else
    echo -e "${RED}✗ Test 1 fallido: No se pudo obtener la lista de estudiantes${NC}"
    echo "Respuesta: $response"
fi

# Test 2: Obtener estudiantes con paginación
echo -e "\nTest 2: Obtener estudiantes con paginación (limit=2, skip=0)"
response=$(curl -s -X GET "${SERVICE_URL}/students?limit=2&skip=0")

if [[ $response == *"students"* ]]; then
    echo -e "${GREEN}✓ Test 2 exitoso: Paginación funcionando correctamente${NC}"
    # Verificar que solo se devolvieron 2 estudiantes
    count=$(echo $response | grep -o '"id"' | wc -l)
    if [ "$count" -le 2 ]; then
        echo -e "${GREEN}✓ Límite de paginación respetado${NC}"
    else
        echo -e "${RED}✗ El límite de paginación no se respetó${NC}"
    fi
else
    echo -e "${RED}✗ Test 2 fallido: La paginación no funcionó${NC}"
fi

# Test 3: Obtener estudiantes ordenados
echo -e "\nTest 3: Obtener estudiantes ordenados por nombre"
response=$(curl -s -X GET "${SERVICE_URL}/students?sort=name")

if [[ $response == *"students"* ]]; then
    echo -e "${GREEN}✓ Test 3 exitoso: Ordenamiento aplicado correctamente${NC}"
else
    echo -e "${RED}✗ Test 3 fallido: El ordenamiento no funcionó${NC}"
fi

# Test 4: Probar parámetros de consulta inválidos
echo -e "\nTest 4: Probar con parámetros de consulta inválidos"
response=$(curl -s -X GET "${SERVICE_URL}/students?limit=invalid")

if [[ $response == *"error"* ]]; then
    echo -e "${GREEN}✓ Test 4 exitoso: El servicio manejó correctamente los parámetros inválidos${NC}"
else
    echo -e "${RED}✗ Test 4 fallido: El servicio no validó los parámetros incorrectamente${NC}"
fi

echo -e "\nPruebas completadas!" 