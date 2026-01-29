const server = 'http://localhost:8080';

async function getTasks(subject){
    try {
        const response = await fetch(server + '/getTasks?subject=' + encodeURIComponent(subject),{
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

document.addEventListener('DOMContentLoaded', function() {
    var subjectElement = document.getElementById("subject");
    subjectElement.addEventListener("change", async function(){
        var selectedValue = subjectElement.value;
        if (selectedValue === "none") return;

        var tasksDiv = document.querySelector(".tasks");
        try {
            const tasks = await getTasks(selectedValue);
            if (tasks === null){
                tasksDiv.innerHTML = "Error";
                return;
            }
            let htmlContent = '';
            tasks.forEach((task, index) => {
                htmlContent += `<div class="task-item">
                    <h3>Задание ${index + 1}</h3>
                    <p>${task.task.replace(/\n/g, '<br>')}</p>
                    <p><strong>Ответ:</strong> ${task.answer}</p>
                </div>`;
            });
            tasksDiv.innerHTML = htmlContent;
        } catch {
            tasksDiv.innerHTML = "Error";
        }
    });
});