$.ajaxSetup({
    cache:false
});

(function($) {
    $.fn.load_students = function() {
        $.ajax({
            url: "/ajax", success: function(result) {
                console.log(result)
            }
        });
    }
})(jQuery);