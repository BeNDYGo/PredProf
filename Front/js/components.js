server = 'http://localhost:8080'

document.getElementById('header-container').innerHTML = `
        <div class='header'>
        <div class='logo' onclick='goToMain()'>TEST</div>
        <div class='navigation'>
            <button class='headers-btn' onclick='goToTasks()'>Задания</button>
            <button class='headers-btn' onclick='goToPVP()'>PVP</button>
        </div>
        <div id='auth-section'>
            <button class='headers-btn' onclick='goToRegister()'>Регистрация</button>
            <button class='headers-btn' onclick='login()'>Войти</button>
        </div>
        </div>
`

function goToMain() {
    window.location.href = 'main.html'
}
function goToPVP() {
    window.location.href = 'pvp.html'
}
function goToAdmin() {
    window.location.href = 'admin.html'
}
async function checkAdmin() {
    const username = localStorage.getItem('username')
    if (username) {
        const url = server + '/api/userInfo?username=' + username
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            }  
        })
        const data = await response.json()
        console.log(data.role)
        if (data.role === 'admin') {
            document.getElementById('auth-section').innerHTML += `
                <button class='headers-btn' onclick='goToAdmin()'>Админ панель</button>
            `;
        }
    }
}

document.addEventListener('DOMContentLoaded', checkAdmin)
