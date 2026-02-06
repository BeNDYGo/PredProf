server = 'http://localhost:8080'

async function getUserInfo(username) {
    const currentUser = localStorage.getItem('username')
    if (!currentUser) {
        return { error: 'not authenticated' }
    }
    
    url = server + '/api/getUserAllInfo?username=' + username
    respons = await fetch(url, {
        headers: {
            'X-Username': currentUser
        }
    })
    const data = await respons.json()
    console.log(data)
    return data
}

const checkButton = document.getElementById('check')
const usernameInput = document.getElementById('username-input')
const lableuserInfo = document.getElementById('user-info')

checkButton.addEventListener('click', async () => {
    username = usernameInput.value
    info = await getUserInfo(username)
    lableuserInfo.innerHTML = ''
    
    if (info.error) {
        lableuserInfo.innerHTML = `<div class="error">Ошибка: ${info.error}</div>`
        return
    }
    
    lableuserInfo.innerHTML = `
        <div class="user-info-card">
            <h3>Информация о пользователе</h3>
            <p><strong>Имя пользователя:</strong> ${info.username}</p>
            <p><strong>Email:</strong> ${info.email}</p>
            <p><strong>Роль:</strong> ${info.role}</p>
            <p><strong>Рейтинг:</strong> ${info.rating}</p>
        </div>
    `
})
