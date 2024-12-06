$.ajaxSetup({
    cache:false
});

(function($) {
    $.fn.load_students = function() {
        $.ajax({
            url: "/ajax", success: function(result) {
                var data = $.parseJSON(result)

                $.each(data["Clients"], function(i, item) {
                    var tbl_body = document.createElement("tbody");
                    $.each(item, function(attr) {
                        // holy shit
                        var row = tbl_body.insertRow();
                        var cell_name = row.insertCell();
                        cell_name.appendChild(document.createTextNode(attr.toString()));
                        var cell_info = row.insertCell();
                        cell_info.appendChild(document.createTextNode(item[attr].toString()));
                    });
                    $("#table").append(tbl_body);
                });
                $.each(data["ClientsTwo"], function(i, item) {
                    var tbl_body = document.createElement("tbody");
                    $.each(item, function(attr) {
                        // holy shit
                        var row = tbl_body.insertRow();
                        var cell_name = row.insertCell();
                        cell_name.appendChild(document.createTextNode(attr.toString()));
                        var cell_info = row.insertCell();
                        cell_info.appendChild(document.createTextNode(item[attr].toString()));
                    });
                    $("#table-two").append(tbl_body);
                });
            }
        });
        return this;
    };
    $.fn.send_students = function() {
        console.log("pong")
    }

})(jQuery);