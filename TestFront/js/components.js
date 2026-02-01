document.getElementById('header-container').innerHTML = `
        <div class="header">
        <div class="logo" onclick="goToMain()">TEST</div>
        <div class="navigation">
            <button class="tasks-btn" onclick="goToTasks()">Задания</button>
        </div>
        <div id="auth-section">
            <button class="register-btn" onclick="goToRegister()">Регистрация</button>
            <button class="login-btn" onclick="login()">Войти</button>
        </div>
        </div>
`

function goToMain() {
    window.location.href = 'main.html';
}