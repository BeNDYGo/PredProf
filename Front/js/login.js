const server = 'https://deck-bedroom-peace-maximum.trycloudflare.com';

document.getElementById('loginForm').addEventListener('submit', async function(e) {
    e.preventDefault();
    
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const messageDiv = document.getElementById('message');
    
    try {
        const response = await fetch(server + '/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                password: password
            })
        });
        
        const result = await response.json();
        
        if (response.ok) {
            localStorage.setItem('username', username);
            window.location.href = 'main.html';
        } else {
            messageDiv.textContent = result.error;
            messageDiv.style.color = 'red';
        }
    } catch (error) {
        messageDiv.textContent = 'Error: ' + error.message;
        messageDiv.style.color = 'red';
    }
});