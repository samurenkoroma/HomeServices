curl -X POST http://localhost:8080/commands/ -H "Content-Type: application/json" \
  -d '{
    "command": "CreateFacility",
    "data": {
      "id": "781643dc-7080-4d24-b3ab-84767d7a1217",
      "name": "Теплица 1",
      "type": "GREENHOUSE",
      "length": 100,
      "width": 50
    }
  }'

# 2. Добавляем грядку в теплицу
curl -X POST http://localhost:8080/commands/ -H "Content-Type: application/json" \
  -d '{
    "command": "AddBed",
    "data": {
      "facility_id": "781643dc-7080-4d24-b3ab-84767d7a1217",
      "name": "Грядка A1",
      "length": 10,
      "width": 1.2
    }
  }'

# 3. Проверяем результат
curl -X POST http://localhost:8080/queries/ -H "Content-Type: application/json" \
  -d '{
    "query": "GetFacilityOverview",
    "data": {
      "facility_id": "greenhouse-123"
    }
  }