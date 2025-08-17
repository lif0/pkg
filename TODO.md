# common

- [ ] Настроить git-workflow(автозапуск тестов, бейдж покрытия тестами)
- [ ] Выпустить первую версию v1.0.0

- [ ]Бейджи как тут https://github.com/kshard/chatter/tree/456d9ae64dea7d70b774b857ab8377032509397b?tab=readme-ov-file

- [] Бейджи
У Coveralls есть «Subprojects» (concurrency, semantic), но отдельные публичные бейджи для них он не генерирует — на странице каждого сабпроекта показан тот же общий бейдж репозитория. 
Coveralls.io
+1

Если всё-таки хочется два раздельных бейджа по модулям, можно сделать кастомные через shields.io (dynamic endpoint) и GitHub Actions, которые будут читать покрытие и публиковать JSON — подскажу, как собрать, когда решишься.

https://img.shields.io/badge/coverage-99%25-green

# semantic

## EstimatePayloadOf
- [ ] Попробовать превысчитать размеры для кейсов где используется unsafe и сравнить разницу

```bash
go test -bench=ArrayValue -benchmem > old.txt
go test -bench=ArrayPointer -benchmem > new.txt
benchstat old.txt new.txt
```