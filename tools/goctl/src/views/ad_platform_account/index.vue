<template>
    <ad-platform-account-editor-form
        v-loading="loading"
        ref="refEditorForm"
        mode="page"
        :is-edit="true"
        :adPlatformAccount="adPlatformAccountItem"
        @submit="onSubmit"
    />
</template>

<script setup lang="ts">
    import { ref, computed, onMounted, onUnmounted } from 'vue';
    import { useRoute } from 'vue-router';
    import { useI18n } from 'vue-i18n';

    import { AdPlatformAccountEditorForm, type AdPlatformAccountEditorFormInstance } from './form';

    import { fetchAdPlatformAccountItem, fetchUpdateAdPlatformAccount, type AdPlatformAccountItem, type ItemReq } from '@/api/ad-platform-account';

    import { message } from '@/utils/message';

    defineOptions({ name: 'AdPlatformAccountEditorPage' });

    const refEditorForm = ref<AdPlatformAccountEditorFormInstance>(null);

    const { t } = useI18n();
    const route = useRoute();

    const adPlatformAccountId = computed(() => Number(route.params.id));

    const loading = ref<boolean>(false);
    const adPlatformAccountItem = ref<AdPlatformAccountItem>({});

    const controller = new AbortController();

    const loadAdPlatformAccountItemData = (id: number) => {
        if (loading.value) controller.abort();

        loading.value = true;
        const data: ItemReq = { id };
        fetchAdPlatformAccountItem({ data, signal: controller.signal })
            .then(
                res => {
                    adPlatformAccountItem.value = res.data;
                },
                err => {
                    if (!err?.isCancelRequest) {
                        message(t('common.dataLoadFailed', { error: String(err) }), { type: 'error' });
                    }
                }
            )
            .finally(() => {
                loading.value = false;
            });
    };

    onMounted(() => loadAdPlatformAccountItemData(adPlatformAccountId.value));

    onUnmounted(() => {
        if (loading.value) {
            controller.abort();
            loading.value = false;
        }
    });

    const onSubmit = () => {
        refEditorForm.value?.validate((valid, _fields) => {
            if (valid) {
                doUpdateConfig(refEditorForm.value.formData);
            }
        });
    };

    const doUpdateAdPlatformAccount = (info: AdPlatformAccountItem) => {
        const data = {
            id: adPlatformAccountId.value,
            name: info.name,
            data: info.data,
            code: info.code,
            type_id: info.type_id,
            type_name: info.type_name,
            type_code: info.type_code,
            schema: info.schema,
            remark: info.remark,
            status: info.status
        };
        fetchUpdateAdPlatformAccount({ data }).then(
            () => {
                message(t('common.changeStatusMessage', { action: t('common.update'), content: info.name }), {
                    type: 'success'
                });
            },
            err => {
                message(
                    t('common.changeStatusFailed', {
                        action: t('common.update'),
                        content: info.name,
                        msg: err
                    }),
                    { type: 'error' }
                );
            }
        );
    };
</script>

<style lang="scss" scoped></style>
