import { useEffect, useState } from "react";

const API_URL = "http://localhost:8000/users/";

interface User {
    id: number;
    name: string;
    email: string;
}

function PageTest() {
    const [users, setUsers] = useState<User[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        // Fazendo a requisição diretamente no useEffect
        const loadUsers = async () => {
            try {
                const response = await fetch(API_URL);

                if (!response.ok) {
                    throw new Error(`Erro na requisição: ${response.status}`);
                }

                const data = await response.json();
                setUsers(data.users); // Aqui esperamos que a resposta tenha { users: User[] }
            } catch (err) {
                setError("Erro ao carregar usuários");
            } finally {
                setLoading(false);
            }
        };

        loadUsers(); // Chamando a função que faz a requisição
    }, []); // Array vazio para rodar uma única vez

    if (loading) return <p>Carregando...</p>;
    if (error) return <p>{error}</p>;

    return (
        <div>
            <h1>Usuários</h1>
            <ul>
                {users.map((user) => (
                    <li key={user.id}>
                        {user.name} - {user.email}
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default PageTest;
