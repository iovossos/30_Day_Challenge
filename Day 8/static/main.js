document.addEventListener("DOMContentLoaded", function () {
    // Handle Signup
    document.getElementById("signupForm")?.addEventListener("submit", function (e) {
        e.preventDefault();

        const username = document.getElementById("signupUsername").value;
        const password = document.getElementById("signupPassword").value;

        fetch('/signup', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password })
        }).then(response => {
            if (response.ok) {
                document.getElementById("signupMessage").textContent = 'Signup successful! Redirecting to login...';
                setTimeout(() => {
                    window.location.href = "/static/login.html";
                }, 1500);
            } else {
                response.text().then((text) => {
                    document.getElementById("signupMessage").textContent = 'Signup failed: ' + text;
                });
            }
        }).catch(error => {
            document.getElementById("signupMessage").textContent = 'Error: ' + error;
        });
    });

    // Handle Login
    document.getElementById("loginForm")?.addEventListener("submit", function (e) {
        e.preventDefault();
        const username = document.getElementById("loginUsername").value;
        const password = document.getElementById("loginPassword").value;

        fetch('/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password })
        }).then(response => {
            if (response.ok) {
                document.getElementById("loginMessage").textContent = 'Login successful! Redirecting to dashboard...';
                setTimeout(() => {
                    window.location.href = "/static/dashboard.html"; // Correct path
                }, 1500);
            } else {
                response.text().then((text) => {
                    document.getElementById("loginMessage").textContent = 'Login failed: ' + text;
                });
            }
        });
    });

    // Handle Task Creation
    document.getElementById("taskForm")?.addEventListener("submit", function (e) {
        e.preventDefault();
        const title = document.getElementById("taskTitle").value;
        const content = document.getElementById("taskContent").value;

        fetch('/api/tasks', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ title, content })
        }).then(response => {
            if (response.ok) {
                alert('Task created successfully!');
                listTasks(); // Refresh the task list after creating a new task
            } else {
                alert('Failed to create task, please log in.');
            }
        });
    });

    // Handle Task Listing
    document.getElementById("listTasks")?.addEventListener("click", function () {
        listTasks(); // Fetch and display tasks when the button is clicked
    });

    function listTasks() {
        fetch('/api/tasks').then(response => response.json()).then(tasks => {
            const taskList = document.getElementById("taskList");
            taskList.innerHTML = '';  // Clear previous tasks

            tasks.forEach(task => {
                const li = document.createElement("li");
                li.textContent = `ID: ${task.id}, Title: ${task.title}, Content: ${task.content}`;
                taskList.appendChild(li);
            });
        }).catch(error => {
            console.error('Error fetching tasks:', error);
        });
    }

    // Handle Logout
    document.getElementById("logoutButton")?.addEventListener("click", function () {
        document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;"; // Clear the token cookie
        window.location.href = "static/login.html"; // Redirect to login page
    });
});
