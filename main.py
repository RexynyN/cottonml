import fastmath
import time
import random


print(fastmath)

# Gerando dados massivos para teste
data = [random.random() * 100 for _ in range(1_000_000)]

print("Processando 1 milhão de números em C...")
start = time.time()

# Chamando nossa função C
result = fastmath.analyze(data)

end = time.time()

print(f"Resultado: {result}")
print(f"Tempo de execução: {end - start:.4f} segundos")

# Teste de erro proposital
try:
    fastmath.analyze(["texto", 1])
except TypeError as e:
    print(f"\nErro capturado com sucesso: {e}")