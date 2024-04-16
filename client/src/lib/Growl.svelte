<script lang="ts">
  import { Icon, type IconType } from 'svelte-icons-pack';
  import { FaSolidCircleInfo } from 'svelte-icons-pack/fa';
  import { FaSolidCircleCheck } from 'svelte-icons-pack/fa';
  import { FaSolidTriangleExclamation } from 'svelte-icons-pack/fa';
  import { FaSolidCircleXmark } from 'svelte-icons-pack/fa';

  type GrowlType = 'info' | 'warning' | 'error' | 'success';
  
  let message: string = 'Notification';
  let growlType: GrowlType = 'info';
  let isShow: boolean = false;
  let icon: IconType = FaSolidCircleInfo;
  
  function show( msg: string, typ: GrowlType, ico: IconType) {
    icon = ico, growlType = typ, isShow = true, message = msg;
    setTimeout( () => isShow = false, 3000 );
  }
  
  export const ShowInfo = function( msg: string ): void {
    show( msg, 'info', FaSolidCircleInfo );
  };
  
  export const ShowWarning = function( msg: string ): void {
    show( msg, 'warning', FaSolidTriangleExclamation );
  };
  
  export const ShowError = function( msg: string ): void {
    show( msg, 'error', FaSolidCircleXmark );
  };
  
  export const ShowSuccess = function( msg: string ): void {
    show( msg, 'success', FaSolidCircleCheck );
  };
</script>

{#if isShow}
  <div class={`growl ${growlType}`}>
    <Icon className='icon_growl' size='20' color='#FFF' src={icon} />
    <span>{message}</span>
  </div>
{/if}

<style>
  .growl.hidden {
    display : none;
  }

  .growl {
    position : fixed;
    display : flex;
    flex-direction : row;
    align-items : center;
    gap : 10px;
    font-size : 14px;
    bottom : 20px;
    right : 20px;
    padding : 10px 20px;
    border-radius : 3px;
    box-shadow : 0px 4px 24px 0px rgba(0, 0, 0, 0.25);
    z-index : 9999;
    min-width : 120px;
    max-width : 350px;
    height : fit-content;
  }

  :global(.icon_growl) {
    flex-shrink : 0;
  }

  .growl span {
    flex-shrink : 1;
  }

  .growl.info {
    background-color : #1080E8;
    color : #FFF;
  }

  .growl.success {
    background-color : #059669;
    color : #FFF;
  }

  .growl.error {
    background-color : #EF4444;
    color : #FFF;
  }

  .growl.warning {
    background-color : #D97706;
    color : #FFF;
  }
</style>