function goToRegister() {
    window.location.href = 'register.html';
}

function goToTasks() {
    window.location.href = 'tasks.html';
}

function checkAuth() {
    const username = localStorage.getItem('username');
    if (username) {
        document.getElementById('auth-section').innerHTML = `
            <span class='username'>${username}</span>
            <button class='headers-btn' onclick='logout()'>Выйти</button>
        `;
    }
}

function logout() {
    localStorage.removeItem('username');
    location.reload();
}

function login() {
    window.location.href = 'login.html';
}

window.onload = checkAuth;
