const server = 'http://localhost:8080';
const wsServer = 'ws://localhost:8080';
const username = localStorage.getItem('username');
const userLableInfo = document.getElementById("userInfo");

async function getUserInfo(username) {
    url = server + '/api/userInfo?username=' + username;
    const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            }
        });
    
    if (response.ok) {
        const data = await response.json();
        return data;
    } else {
        return null
    }
}

async function eloRender(){
    const data = await getUserInfo(username);

    if (data) {
        userLableInfo.innerHTML += `
            <div>
                ${data.username}: ${data.rating} ELO
            </div>
        `;
    }
}

eloRender()

function sendAnswer() {
    const answerInput = document.getElementById('answerInput');
    const answer = answerInput.value.trim();
    
    if (answer && socket) {
        socket.send(JSON.stringify({"userAnswer": answer}));
        answerInput.value = '';
    }
}

function searchGame(){
    socket = new WebSocket(wsServer + '/api/ws');
    socket.onopen = () => {
        console.log("Connected");
    };
    socket.onmessage = (event) => {
        const data = event.data;
        const parsed = JSON.parse(data);
        const taskDiv = document.getElementById('divTask');
        
        if (parsed.task) {
            
            taskDiv.innerHTML = `<div class="task-container">
                <div class="task-text">${parsed.task.replace(/\n/g, '<br>')}</div>
                </div>`;
        } else {
            const msg = document.getElementById("messageContainer");
            console.log(parsed.message);
            
            let messageText = "";
            if (parsed.message === "match found") {
                messageText = "Матч найден!";
                msg.style.color = "green";
            } else if (parsed.message === "waiting opponent...") {
                messageText = "Ожидание противника...";
                msg.style.color = "orange";
            } else if (parsed.message === "opponent disconected") {
                messageText = "Противник отключился!";
                msg.style.color = "red";
                taskDiv.textContent = "";
            } else if (parsed.message === "incorrect"){
                messageText = "Неверный ответ!";
                msg.style.color = "red";
            } else if (parsed.message === "correct"){
                messageText = "Верный ответ!";
                msg.style.color = "green";
                taskDiv.textContent = "";
            } else if (parsed.message === "you win"){
                messageText = "Вы победили!";
                msg.style.color = "green";
                taskDiv.textContent = "";
            } else if (parsed.message === "you lose"){
                messageText = "Противник победил!";
                msg.style.color = "red";
                taskDiv.textContent = "";
            }
            
            if (messageText) {
                msg.innerHTML = `<div style="color: ${msg.style.color}; margin: 2px 0;">${messageText}</div>` + msg.innerHTML;
            }
        }
    }
}
