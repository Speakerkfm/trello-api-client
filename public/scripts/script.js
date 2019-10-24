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


function createCard(cardName, cardSecription) {
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