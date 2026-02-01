const server = 'http://localhost:8080';

async function getTasks(subject, taskType = "", difficulty = ""){
    try {
        let url = server + '/getTasks?subject=' + encodeURIComponent(subject);
        if (taskType && taskType !== "none") {
            url += '&taskType=' + encodeURIComponent(taskType);
        }
        if (difficulty && difficulty !== "none") {
            url += '&difficulty=' + encodeURIComponent(difficulty);
        }
        
        const response = await fetch(url, {
            method: "GET",
            headers: {
                'Content-Type': 'application/json',
            }
        });

        if (response.ok){
            return await response.json();
        } else {
            return null;
        }
    } catch (error){
        return null;
    }
}

function showAnswer(index){
    var answerElement = document.getElementById(`answer-${index}`);
    answerElement.style.display = answerElement.style.display === 'none' ? 'block' : 'none';
}

document.addEventListener('DOMContentLoaded', async function() {
    var subjectElement = document.getElementById("subject");
    var taskTypeElement = document.getElementById("taskType");
    var difficultyElement = document.getElementById("difficulty");
    
    async function loadTasks() {
        var selectedSubject = subjectElement.value;
        var selectedTaskType = taskTypeElement.value;
        var selectedDifficulty = difficultyElement.value;
        
        if (selectedSubject === "none") return;

        var tasksDiv = document.getElementById("tasks");
        try {
            const tasks = await getTasks(selectedSubject, selectedTaskType, selectedDifficulty);
            if (tasks === null){
                tasksDiv.innerHTML = "Нет заданий";
                return;
            }
            let htmlContent = '';
            tasks.forEach((task, index) => {
                htmlContent += `<div class="task-item">
                    <h3>----------</h3>
                    <p>${task.task.replace(/\n/g, '<br>')}</p>
                    <button class="show-answer-btn" onclick="showAnswer(${index})">Показать ответ</button>
                    <p id="answer-${index}" style="display: none"><strong>Ответ:</strong> ${task.answer}</p>
                </div>`;
            });
            tasksDiv.innerHTML = htmlContent;
        } catch (error) {
            tasksDiv.innerHTML = "Error";
        }
    }
    
    subjectElement.addEventListener("change", loadTasks);
    taskTypeElement.addEventListener("change", loadTasks);
    difficultyElement.addEventListener("change", loadTasks);
});
