async function getTasks(subject){
    try {
        const response = await fetch('http://localhost:8080/tasks',{
            method: "GET",
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                subject: subject
            })
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

var subjectElement = document.getElementById("subject");
subjectElement.addEventListener("change", async function(){
    var selectedValue = subjectElement.value;
    if (selectedValue === "none") return;

    var tasksDiv = document.getElementById("tasks");
    try {
        const tasks = await getTasks(selectedValue);
        if (tasks === null){
            tasksDiv.innerHTML = "Error";
            return;
        }
        tasksDiv.innerHTML = tasks;
    } catch {
        tasksDiv.innerHTML = "Error";
    }
});