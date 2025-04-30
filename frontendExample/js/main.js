const templates = {
    'feat-interface': {status: 'feat', title: '新增 API interface'},
    'feat-impl': {status: 'feat', title: '實作 API'},
    'test': {status: 'test', title: '新增單元測試'},
    'docs': {status: 'docs', title: '新增 swagger 註解'}
};

let commitHistory = JSON.parse(localStorage.getItem('commitHistory')) || [];

$(document).ready(function () {
    $('#templateSelect').on('change', function () {
        const template = templates[this.value];
        if (template) {
            $('#status').val(template.status);
            $('#title').val(template.title);
        }
    });

    $('#commit-form').on('submit', function (event) {
        event.preventDefault();

        const status = $('#status').val();
        const title = $('#title').val();
        const apiGroup = $('#group').val();
        const route = $('#route').val();
        const method = $('#method').val();
        const description = $('#description').val();
        const issueContent = $('#issueContent').val();

        const APIGroups = apiGroup.split("/");
        let formattedRoute = route.replace(/\/:(\w+)/g, '/{$1}');
        let formattedFullPathStr = "";
        for (let i = 0; i < APIGroups.length; i++) {
            formattedFullPathStr += `/api/${APIGroups[i].toLowerCase()}${formattedRoute}\n`
        }

        const commitTitle = `${status}: ${title} 🔗${method} [${apiGroup}]${route}`;
        const commitMessage = `- Feature: ${description}\n- API Route: ${method}\n${formattedFullPathStr}\n${issueContent}`;

        $('#commitTitle').text(commitTitle);
        $('#commitMessage').text(commitMessage);

        const historyItem = {
            timestamp: new Date().toISOString(),
            title: commitTitle,
            message: commitMessage,
            template: $('#templateSelect').val(),
            inputs: {
                status, title, apiGroup, route, method, description, issueContent
            }
        };
        commitHistory.unshift(historyItem);
        if (commitHistory.length > 15) {
            commitHistory.pop();
        }

        localStorage.setItem('commitHistory', JSON.stringify(commitHistory));

        updateRecentHistoryList();
        updateCompleteHistoryList();
    });

    $('.copy-btn').on('click', function () {
        const targetId = $(this).data('target');
        const text = $('#' + targetId).text();

        navigator.clipboard.writeText(text).then(() => {
            const originalText = $(this).text();
            $(this).text('Copied!');
            setTimeout(() => {
                $(this).text(originalText);
            }, 2000);
        });
    });

    $('#clearHistoryBtn').click(function () {
        commitHistory = [];
        localStorage.removeItem('commitHistory');

        updateRecentHistoryList();
        updateCompleteHistoryList();
    });

    function updateRecentHistoryList() {
        const historyList = $('#recentHistoryList');
        historyList.empty();

        const recentHistory = commitHistory.slice(0, 3);

        recentHistory.forEach((item, index) => {
            const historyItem = createHistoryItem(item, index);
            historyList.append(historyItem);
        });

        if (recentHistory.length === 0) {
            historyList.text('No recent history found.');
        }

        if (commitHistory.length > 3) {
            $('#viewAllBtn').show();
        } else {
            $('#viewAllBtn').hide();
        }
    }

    function updateCompleteHistoryList() {
        const historyList = $('#completeHistoryList');
        historyList.empty();

        commitHistory.forEach((item, index) => {
            const historyItem = createHistoryItem(item, index);
            historyList.append(historyItem);
        });

        if (commitHistory.length === 0) {
            historyList.text('No complete history found.');
        }

    }

    function createHistoryItem(item, index) {
        const historyItem = $(`
                    <div class="history-item" data-index="${index}">
                        <div class="history-timestamp">${new Date(item.timestamp).toLocaleString()}</div>
                        <div>${item.title}</div>
                    </div>
                `);

        historyItem.click(function () {
            const itemIndex = $(this).data('index');
            const selectedItem = commitHistory[itemIndex];

            $('#status').val(selectedItem.inputs.status);
            $('#title').val(selectedItem.inputs.title);
            $('#group').val(selectedItem.inputs.apiGroup);
            $('#route').val(selectedItem.inputs.route);
            $('#method').val(selectedItem.inputs.method);
            $('#description').val(selectedItem.inputs.description);
            $('#issueContent').val(selectedItem.inputs.issueContent);

            $('#commitTitle').text(selectedItem.title);
            $('#commitMessage').text(selectedItem.message);

            $('#historyModal').hide();
        });

        return historyItem;
    }


    $('#viewAllBtn').click(function () {
        $('#historyModal').show();
        updateCompleteHistoryList();
    });


    $('.close-btn').click(function () {
        $('#historyModal').hide();
    });

    $(window).click(function (event) {
        if ($(event.target).is('#historyModal')) {
            $('#historyModal').hide();
        }
    });


    if (commitHistory.length > 0) {
        updateRecentHistoryList();
    } else {
        $('#recentHistoryList').text('No recent history found.');
    }
});