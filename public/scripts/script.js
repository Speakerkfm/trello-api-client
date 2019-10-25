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
            location.reload();
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

function updateOwnCardStatus(cardID, listID) {
    $.ajax({
        type: "POST",
        data: {
            "listID": listID
        },
        url: "/v2/cards/" + cardID,
        success: function () {
            location.reload();
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