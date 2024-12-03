$.ajaxSetup({
    cache:false
});

(function($) {
    $.fn.load_rating_data = function() { 
        //Get userID to use info for
        var parts = window.location.pathname.split('/');
        var lastSegment = parts.pop() || parts.pop();  //Handle potential trailing slash
        ID = parseInt(lastSegment) - 1;

        $.getJSON("/web/data/users.json", function(data) {
            user = data[ID].username;
            console.log(user);
            $("#username").append(`<h3>` + user + `</h3>`);

            rating_data = data[ID].ratings;
            $.each(rating_data, function(value) {
                //Slightly ew and i noticed why but im keeping it because its funny
                $("#rating-data").append(value + " " + "-" + " " + rating_data[value] + `<br>`);
            });
        });
        return this;
    };
    $.fn.load_all_users = function() {
        $.getJSON("/web/data/users.json", function(data) {
            $.each(data, function(user) {
                $("#main-landing").append(`<a href="/userpage/` + data[user].id + `">` + data[user].username + `</a><br><br>`);
            });
        });

        return this;
    }
})(jQuery);