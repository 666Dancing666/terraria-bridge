package com.terrariabridge.gui;

import com.terrariabridge.network.BridgeClient;
import net.minecraft.client.Minecraft;
import net.minecraft.client.gui.GuiGraphics;
import net.minecraft.client.gui.screens.Screen;
import net.minecraft.client.gui.components.Button;
import net.minecraft.network.chat.Component;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.HashMap;

public class RecipeScreen extends Screen
{
    private List<Map<String, Object>> recipes = new ArrayList<>();
    private int scrollOffset = 0;
    private static final int ITEMS_PER_PAGE = 8;
    private BridgeClient bridgeClient;

    public RecipeScreen(BridgeClient bridgeClient)
    {
        super(Component.literal("Terraria Crafting"));
        this.bridgeClient = bridgeClient;
    }

    public void setRecipes(List<Map<String, Object>> recipes)
    {
        this.recipes = recipes;
        this.scrollOffset = 0;
    }

    @Override
    protected void init()
    {
        super.init();
        rebuildButtons();
    }

    private void rebuildButtons()
    {
        this.clearWidgets();
        int startX = this.width / 2 - 100;
        int startY = 40;

        int end = Math.min(scrollOffset + ITEMS_PER_PAGE, recipes.size());
        for (int i = scrollOffset; i < end; i++)
        {
            final Map<String, Object> recipe = recipes.get(i);
            final String key = (String) recipe.get("key");
            String name = (String) recipe.get("result_name");
            int count = ((Number) recipe.get("result_count")).intValue();
            String label = name + " x" + count;
            int y = startY + (i - scrollOffset) * 22;

            this.addRenderableWidget(Button.builder(
                Component.literal(label),
                btn -> {
                    if (bridgeClient != null)
                    {
                        Map<String, Object> payload = new HashMap<>();
                        payload.put("recipe", key);
                        bridgeClient.send(
                            com.terrariabridge.network.MessageTypes.toJson("craft_item", payload));
                    }
                    Minecraft mc = Minecraft.getInstance();
                    if (mc.player != null)
                    {
                        mc.player.displayClientMessage(
                            Component.literal("Crafted: " + name), false);
                    }
                })
                .pos(startX, y)
                .size(200, 20)
                .build());
        }

        if (recipes.size() > ITEMS_PER_PAGE)
        {
            if (scrollOffset > 0)
            {
                this.addRenderableWidget(Button.builder(
                    Component.literal("^ Up"),
                    btn -> { scrollOffset = Math.max(0, scrollOffset - ITEMS_PER_PAGE); rebuildButtons(); })
                    .pos(startX, startY - 22)
                    .size(200, 20)
                    .build());
            }
            if (scrollOffset + ITEMS_PER_PAGE < recipes.size())
            {
                this.addRenderableWidget(Button.builder(
                    Component.literal("v Down"),
                    btn -> { scrollOffset += ITEMS_PER_PAGE; rebuildButtons(); })
                    .pos(startX, startY + ITEMS_PER_PAGE * 22)
                    .size(200, 20)
                    .build());
            }
        }
    }

    @Override
    public void render(GuiGraphics graphics, int mouseX, int mouseY, float partialTick)
    {
        this.renderBackground(graphics);
        graphics.drawCenteredString(this.font, "Terraria Crafting (R to close)", this.width / 2, 15, 0xFFFFFF);
        super.render(graphics, mouseX, mouseY, partialTick);

        int startY = 40;
        int end = Math.min(scrollOffset + ITEMS_PER_PAGE, recipes.size());
        for (int i = scrollOffset; i < end; i++)
        {
            Map<String, Object> recipe = recipes.get(i);
            String station = (String) recipe.get("station");
            int y = startY + (i - scrollOffset) * 22 + 16;
            graphics.drawString(this.font, "Station: " + station, 20, y, 0xAAAAAA);
        }
    }

    @Override
    public boolean isPauseScreen()
    {
        return false;
    }
}
