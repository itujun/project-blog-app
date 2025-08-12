'use client';
import { Button } from '@/components/ui/button';
import { useQuery } from '@tanstack/react-query';

// fetchHealth memanggil API backend /health
async function fetchHealth() {
  const res = await fetch('http://localhost:8080/health');
  if (!res.ok) throw new Error('Gagal fetch health');
  return res.json() as Promise<{ status: string }>;
}

export default function Home() {
  // useQuery: otomatis cache & refresh sesuai best practice v5
  const { data, isLoading, error } = useQuery({
    queryKey: ['health'],
    queryFn: fetchHealth,
  });
  return (
    <main className="p-6 space-y-4">
      <h1 className="text-2xl font-bold">BlogApp</h1>

      <div>
        <Button variant="default">Contoh Button (shadcn/ui)</Button>
      </div>

      <div className="rounded-lg border p-4">
        {isLoading && <p>Memuat...</p>}
        {error && (
          <p className="text-red-600">Error: {(error as Error).message}</p>
        )}
        {data && (
          <p>
            Backend status: <strong>{data.status}</strong>
          </p>
        )}
      </div>
    </main>
  );
}
