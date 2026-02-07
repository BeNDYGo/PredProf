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
            <p><strong>Победы:</strong> ${info.wins}</p>
            <p><strong>Поражения:</strong> ${info.losses}</p>
        </div>
    `
})

// Добавление задания

const addTaskBtn = document.getElementById('add-task-btn')
const taskStatus = document.getElementById('task-status')

addTaskBtn.addEventListener('click', async () => {
    const currentUser = localStorage.getItem('username')
    if (!currentUser) {
        taskStatus.innerHTML = '<div class="error">Вы не авторизованы</div>'
        return
    }

    const taskName = document.getElementById('task-name').value.trim()
    const taskText = document.getElementById('task-text').value.trim()
    const answer = document.getElementById('task-answer').value.trim()
    const subject = document.getElementById('task-subject').value
    const taskType = document.getElementById('task-type').value
    const difficulty = document.getElementById('task-difficulty').value

    if (!taskName || !taskText || !answer) {
        taskStatus.innerHTML = '<div class="error">Заполните все поля</div>'
        return
    }

    const fullTask = taskName + '\n\n' + taskText

    try {
        const response = await fetch(server + '/api/addTask?subject=' + subject, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'X-Username': currentUser
            },
            body: JSON.stringify({
                task: fullTask,
                answer: answer,
                taskType: taskType,
                difficulty: difficulty
            })
        })

        const data = await response.json()

        if (data.error) {
            taskStatus.innerHTML = `<div class="error">Ошибка: ${data.error}</div>`
        } else {
            taskStatus.innerHTML = '<div class="success">Задание добавлено</div>'
            document.getElementById('task-name').value = ''
            document.getElementById('task-text').value = ''
            document.getElementById('task-answer').value = ''
        }
    } catch (err) {
        taskStatus.innerHTML = `<div class="error">Ошибка: ${err.message}</div>`
    }
})
