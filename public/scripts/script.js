function deleteCardById(id) {
    if (!confirm('Вы уверены?')) {
        return;
    }

    $.ajax({
        type: "DELETE",
        url: "/v1/cards/" + id,
        success: function () {
            location.reload();
        }
    });
}

function deleteOwnCardById(id) {
    if (!confirm('Вы уверены?')) {
        return;
    }

    $.ajax({
        type: "DELETE",
        url: "/v2/cards/" + id,
        success: function () {
            document.querySelector(`li[data-id="${id}"]`).remove();
        }
    });
}

function deleteBoardById(id) {
    if (!confirm('Вы уверены?')) {
        return;
    }

    $.ajax({
        type: "DELETE",
        url: "/v1/boards/" + id,
        success: function () {
            location.reload();
        }
    });
}

function updateCardStatus(cardID, listID) {
    $.ajax({
        type: "POST",
        data: {
            "listID": listID
        },
        url: "/v1/cards/" + cardID,
        success: function () {
            location.reload();
        }
    });
}

function updateOwnCardStatus(cardID, listID, status) {
    $.ajax({
        type: "POST",
        data: {
            "listID": listID
        },
        url: "/v2/cards/" + cardID,
        success: function () {
            document.querySelector(`button[data-id="${cardID}-status"]`).innerHTML = status
        }
    });
}

$('#create-card-form').submit(function(e){
    e.preventDefault();

    $.ajax({
        type: "POST",
        url: '/v1/card',
        data: $(this).serialize(),
        success: function()
        {
            location.reload();
        }
    });
});

$('#create-own-card-form').submit(function(e){

    e.preventDefault();
    if (document.getElementById("task-name").innerText === ""){
        alert("Task name is empty!");
        return
    }

    $.ajax({
        type: "POST",
        url: '/v2/card',
        data: $(this).serialize(),
        success: function()
        {
            location.reload();
        }
    });
});

function loginWithTrello() {
    $.ajax({
        type: "POST",
        url: "/v1/trello/login_redirect" ,
        success: function (data) {
            location.href=data.url;
        }
    });
}

function logout() {
    $.ajax({
        type: "GET",
        url: "/v1/logout" ,
        success: function () {
            location.href="/";
        }
    });
}

const Comments = {
    async fetchComments(cardID) {
        const commentsList = document.getElementById('comment-list');
        const todosResponse = await fetch('https://jsonplaceholder.typicode.com/comments?postId=' + cardID);
        if (!todosResponse.ok) {
            throw new Error('Не удалось получить комментарии... ');
        }
        const comments = await todosResponse.json();

        for (const comment of comments) {
            const commentItem = document.createElement('li');
            // List Item (ToDo)
            commentItem.classList.add('list-group-item');
            commentItem.classList.add('card-item');

            const titleEl = document.createElement('span');
            titleEl.textContent = comment.name;
            titleEl.classList.add('card-desc');
            commentItem.appendChild(titleEl);

            const emailEl = document.createElement('span');
            emailEl.textContent = comment.email;
            emailEl.classList.add('card-desc');
            commentItem.appendChild(emailEl);

            const commentEl = document.createElement('span');
            commentEl.textContent = comment.body;
            commentEl.classList.add('card-desc');
            commentItem.appendChild(commentEl);

            commentsList.appendChild(commentItem);
        }
    },
};