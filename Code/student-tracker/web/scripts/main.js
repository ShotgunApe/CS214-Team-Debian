$.ajaxSetup({
    cache:false
});

(function($) {
    $.fn.load_students = function() {
        $.get("student/student.go", function(data) {
            console.log(data);
        });
        return this;
    }
})(jQuery);