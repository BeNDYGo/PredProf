document.getElementById('header-container').innerHTML = `
        <div class="header">
        <div class="logo" onclick="goToMain()">TEST</div>
        <div class="navigation">
            <button class="headers-btn" onclick="goToTasks()">Задания</button>
            <button class="headers-btn" onclick="goToPVP()">PVP</button>
        </div>
        <div id="auth-section">
            <button class="headers-btn" onclick="goToRegister()">Регистрация</button>
            <button class="headers-btn" onclick="login()">Войти</button>
        </div>
        </div>
`

function goToMain() {
    window.location.href = 'main.html';
}
function goToPVP() {
    window.location.href = 'pvp.html'
}