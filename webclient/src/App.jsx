import { useState } from 'react';

export default function App() {
  const [msg, setMsg] = useState('');

  const loadGreeting = async () => {
    try {
      const res  = await fetch('/api/greeting');
      const data = await res.json();
      setMsg(data.greeting);
    } catch (e) {
      console.error(e);
      setMsg('Error :(');
    }
  };

  return (
    <main style={{ fontFamily: 'system-ui, sans-serif', margin: '2rem' }}>
      <h1>Simple GUI backed by Go + React</h1>
      <button onClick={loadGreeting}>Load greeting</button>
      <p>{msg}</p>
    </main>
  );
}
